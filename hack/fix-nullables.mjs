#!/usr/bin/env node
// Patches an OpenAPI 3.0 spec produced by openapi-down-convert for
// compatibility with oapi-codegen.
//
// Two rewrites are applied (in one pass, idempotent):
//
// 1. Nullable types. openapi-down-convert turns 3.1 `type: ['string', 'null']`
//    into `anyOf: [{type: 'string'}, {type: 'null'}]`. oapi-codegen renders
//    that as an opaque per-field union wrapper struct. The 3.0 form that
//    produces a clean `*string` / `*float32` is `{type: 'string', nullable: true}`.
//    For $ref branches we emit `{allOf: [{$ref: ...}], nullable: true}`.
//
// 2. Enums missing `type`. Zod-to-JSON-Schema emits string enums as
//    `{enum: [...]}` without `type`. OpenAPI 3.0 requires `type` to be set
//    or oapi-codegen falls back to `interface{}`. If every enum value is a
//    string we inject `type: string`.
//
// Usage: node hack/fix-nullables.mjs <path-to-spec.yaml> [out-path]
// Defaults to in-place edit.

import { readFileSync, writeFileSync } from 'node:fs';
import yaml from 'yaml';

const [, , inPath, outPath = inPath] = process.argv;
if (!inPath) {
  console.error('usage: fix-nullables.mjs <spec.yaml> [out]');
  process.exit(2);
}

const doc = yaml.parse(readFileSync(inPath, 'utf8'));

let nullableRewrites = 0;
let enumTypeAdditions = 0;

function isNullBranch(b) {
  return b && typeof b === 'object' && b.type === 'null';
}

function walk(node) {
  if (Array.isArray(node)) {
    node.forEach(walk);
    return;
  }
  if (!node || typeof node !== 'object') return;

  for (const key of ['anyOf', 'oneOf']) {
    const branches = node[key];
    if (!Array.isArray(branches) || branches.length !== 2) continue;
    const nullIdx = branches.findIndex(isNullBranch);
    if (nullIdx === -1) continue;
    const other = branches[1 - nullIdx];
    delete node[key];
    if (other.$ref) {
      node.allOf = [{ $ref: other.$ref }];
      node.nullable = true;
    } else {
      Object.assign(node, other);
      node.nullable = true;
    }
    nullableRewrites++;
  }

  if (Array.isArray(node.enum) && node.type === undefined && node.enum.every((v) => typeof v === 'string')) {
    node.type = 'string';
    enumTypeAdditions++;
  }

  for (const v of Object.values(node)) walk(v);
}

walk(doc);

writeFileSync(outPath, yaml.stringify(doc));
console.error(`fix-spec: nullable=${nullableRewrites} enum-type=${enumTypeAdditions}`);

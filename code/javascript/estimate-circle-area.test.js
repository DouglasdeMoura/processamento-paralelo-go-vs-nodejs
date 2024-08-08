import assert from 'node:assert/strict'
import { test } from 'node:test'
import { estimateCircleArea, estimateCircleAreaParallel } from './estimate-circle-area.js'

test('estimateCircleArea', async () => {
  const result = await estimateCircleArea(1, 1000000)
  assert.ok(result > 3 && result < 3.5)
})

test('estimateCircleAreaParallel', async () => {
  const result = await estimateCircleAreaParallel(1, 1000000)
  assert.ok(result > 3 && result < 3.5)
})


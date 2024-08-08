import Benchmark from 'benchmark'
import { estimateCircleArea, estimateCircleAreaParallel } from './estimate-circle-area.js'

const ray = 1
const samples = 1000000

const suite = new Benchmark.Suite

suite
  .add('estimateCircleArea', (deferred) => {
    estimateCircleArea(ray, samples)
    deferred.resolve()
  }, { defer: true })
  .add('estimateCircleAreaParallel', {
    defer: true,
    fn: async (deferred) => {
      await estimateCircleAreaParallel(ray, samples)
      deferred.resolve()
    }
  })
  .on('cycle', (event) => {
    console.log(String(event.target))
  })
  .on('complete', function () {
    console.log(`\nFastest is ${this.filter('fastest').map('name')}`)
  })
  .run({ 'async': true })
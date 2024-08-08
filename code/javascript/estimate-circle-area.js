import { Worker } from 'node:worker_threads'
import { cpus } from 'node:os'

/**
 * Estimate the area of a circle using Monte Carlo method with multiple workers
 * 
 * @param {number} ray 
 * @param {number} samples 
 * @returns {Promise<number>}
 */
export function estimateCircleAreaParallel(ray, samples) {
  // pick the number of workers based on the number of CPUs
  const workersQuantity = cpus().length

  return new Promise((resolve, reject) => {
    const samplesPerWorker = Math.floor(samples / workersQuantity)
    const workers = []
    let totalInside = 0
    let completedWorkers = 0

    for (let i = 0; i < workersQuantity; i++) {
      const worker = new Worker('./estimate-circle-area-worker.js')
      workers.push(worker)

      worker.on('message', (countInside) => {
        totalInside += countInside
        completedWorkers++
        if (completedWorkers === workersQuantity) {
          workers.forEach(worker => worker.terminate())
          resolve((totalInside / samples) * 4)
        }
      })

      worker.on('error', (err) => { reject(err) })
      worker.postMessage({ ray, samples: samplesPerWorker })
    }
  })
}

/**
 * Estimate the area of a circle using Monte Carlo method
 * 
 * @param {number} ray 
 * @param {number} samples 
 * @returns {number}
 */
export function estimateCircleArea(ray, samples) {
  let pointsInside = 0

  for (let i = 0; i < samples; i++) {
    const x = Math.random()
    const y = Math.random()
    if (x * x + y * y <= ray) {
      pointsInside++
    }
  }

  return (pointsInside / samples) * 4
}

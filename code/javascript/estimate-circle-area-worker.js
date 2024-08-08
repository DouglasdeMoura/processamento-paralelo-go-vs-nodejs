import { parentPort } from 'node:worker_threads'

parentPort.on('message', ({ ray, samples }) => {
  let pointsInside = 0

  for (let i = 0; i < samples; i++) {
    const x = Math.random()
    const y = Math.random()

    if (x * x + y * y <= ray) {
      pointsInside++
    }
  }

  parentPort.postMessage(pointsInside);
})

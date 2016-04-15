import isomorphicFetch from 'isomorphic-fetch'
import URI from 'urijs'

// Will take a relative path and turn it into an absolute path to the API root.
function ensureAbsoluteUrl(apiRoot, path) {
  if (typeof path !== 'string') return path
  if (URI(path).is('absolute')) return path
  return URI(apiRoot + path).normalize().toString()
}

export default function fetch(input, options) {
  const apiRoot = 'http://localhost:3130/'
  const path = ensureAbsoluteUrl(apiRoot, input)

  return isomorphicFetch(path, options).then((r) => r.json())
}

export function records() {
  return fetch(`/records`)
}

export function record(id) {
  return fetch(`/records/${id}`)
}

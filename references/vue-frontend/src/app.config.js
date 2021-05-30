
function fetchEnvironmentVariables() {
  let env

  // If inside a Docker container (environment variables are injected by the entrypoint script)
  if (window.env !== undefined) env = window.env
  // Else ran locally on the developer machine (the environment variables will be imported by the Vite runtime)
  else env = import.meta.env

  return {
    BASE_URL: env.VITE_BASE_URL,
    GATEKEEPER_URL: env.VITE_GATEKEEPER_URL,
    BACKEND_URL: env.VITE_BACKEND_URL,
  }
}

export default fetchEnvironmentVariables()

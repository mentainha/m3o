/**
 * M3O Weather Now
 * -------------------------
 *
 * Get the current weather
 */

import { call } from '..'

async function currentWeather(loc) {
  const response = await call('weather', 'Now', {
    location: loc
  })

  console.log(response)
}


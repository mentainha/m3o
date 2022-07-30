/**
 * M3O Weather Now
 * -------------------------
 *
 * Get the current weather
 */

import { m3oRequest } from '..'

async function currentWeather(loc) {
  const response = await m3oRequest('weather', 'Now', {
    location: loc
  })

  console.log(response)
}


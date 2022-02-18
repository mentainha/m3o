require 'net/http'
require 'json'

uri = URI('https://api.m3o.com/v1/helloworld/Call')
http = Net::HTTP.new(uri.hostname, uri.port)
http.use_ssl = true
req = Net::HTTP::Post.new(uri, 'Content-Type' => 'application/json')
req.body = {name: 'Alice'}.to_json
req['Authorization'] = 'Bearer ' + ENV['M3O_API_TOKEN']

# make request
res = http.request(req)

case res
when Net::HTTPSuccess
  puts res.body
else
  puts res.read_body
end

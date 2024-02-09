url = 'http://httpbin.org/ip'

resp, err = fetch(url, {})

if err ~= nil then
    print('Oh, no!')
else
    json, err  = resp:JSON()

    if err ~= nil then
        fmt.Println('ERROR:', err)
        return
    end

    print(json['origin'])
end

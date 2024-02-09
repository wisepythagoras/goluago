url = 'http://httpbin.org/ip'

function cb(resp, err)
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
end

function my_fetch()
    return fetch(url, {})
end

runAsync(my_fetch, cb)

for i=1,10000000 do
    -- Do nothing
end

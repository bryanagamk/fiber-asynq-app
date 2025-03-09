math.randomseed(os.time())  -- Seed generator untuk nilai acak yang lebih bervariasi

function request()
    local random_number = math.random(1, 1000000000)
    local body = string.format('{"name": "Test User %d", "email": "test%d@example.com"}', random_number, random_number)

    wrk.method = "POST"
    wrk.body = body
    wrk.headers["Content-Type"] = "application/json"

    return wrk.format(nil, nil, nil, body)
end

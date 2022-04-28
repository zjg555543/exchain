res=$(exchaincli tx wasm store ./wasm/escrow/artifacts/cw_escrow-aarch64.wasm --fees 0.01okt --from captain --gas=2000000 -b block -y)
code_id=$(echo "$res" | jq '.logs[0].events[1].attributes[0].value' | sed 's/\"//g')
res=$(exchaincli tx wasm instantiate "$code_id" '{"arbiter":"ex1h0j8x0v9hs4eq6ppgamemfyu4vuvp2sl0q9p3v","end_height":100000,"recipient":"ex190227rqaps5nplhg2tg8hww7slvvquzy0qa0l0"}' --label test1 --admin ex1h0j8x0v9hs4eq6ppgamemfyu4vuvp2sl0q9p3v --fees 0.001okt --from captain -b block -y)
contractAddr=$(echo "$res" | jq '.logs[0].events[0].attributes[0].value' | sed 's/\"//g')
exchaincli tx wasm execute "$contractAddr" '{"approve":{"quantity":[{"amount":"100","denom":"okt"}]}}' --fees 0.001okt --from captain -b block -y
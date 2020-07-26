mkdir -p testdata/releases/10567529/rating
mkdir -p testdata/releases/2000000000/rating
mkdir -p testdata/users/brotherlogic/collection/folders/0/
mkdir -p testdata/users/brotherlogic/collection/folders/812802/releases
mkdir -p testdata/users/brotherlogic/collection/folders/673768/releases/10866041/instances
mkdir -p testdata/users/brotherlogic/collection/releases
mkdir -p testdata/masters/67464
mkdir -p testdata/masters/38998
mkdir -p testdata/masters/38677
mkdir -p testdata/masters/5251
mkdir -p testdata/marketplace/price_suggestion
mkdir -p testdata/marketplace/listings
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/releases/323005?token=$1" | sed "s/$1/token/g" > testdata/users/brotherlogic/collection/releases/323005_t
sleep 1
exit
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/inventory?status=For+Sale&per_page=100&token=$1&page=4" | sed "s/$1/token/g" > testdata/users/brotherlogic/inventory_status=For+Sale_per_page=100_token=token_page=4
sleep 1
exit
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/inventory?status=For+Sale&per_page=100&token=$1&page=3" | sed "s/$1/token/g" > testdata/users/brotherlogic/inventory_status=For+Sale_per_page=100_token=token_page=3
sleep 1
exit
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/inventory?status=For+Sale&per_page=100&token=$1&page=2" | sed "s/$1/token/g" > testdata/users/brotherlogic/inventory_status=For+Sale_per_page=100_token=token_page=2
sleep 1
exit
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/inventory?status=For%20Sale&per_page=100&token=$1" | sed "s/$1/token/g" > testdata/users/brotherlogic/inventory_status=For%20Sale_per_page=100_token=token
sleep 1
exit
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/marketplace/listings/885335204?curr_abbr=USD&token=$1" | sed "s/$1/token/g" > testdata/marketplace/listings/885335204_curr_abbr=USD_token=token
sleep 1
exit
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/marketplace/price_suggestions/1215019?token=$1" |  sed "s/$1/token/g" > testdata/marketplace/price_suggestions/1215019_token=token
sleep 1
exit
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/masters/38677/versions?per_page=100&token=$1&page=2" | sed "s/$1/token/g" > testdata/masters/38677/versions_per_page=100_token=token_page=2
sleep 1
exit
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/masters/38677/versions?per_page=500&token=$1" | sed "s/$1/token/g" > testdata/masters/38677/versions_per_page=500_token=token
sleep 1
exit
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/releases/2535152?token=$1" | sed "s/$1/token/g" > testdata/releases/2535152_token=token
sleep 1
exit
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/marketplace/listings/647415422?curr_abbr=USD&token=$1" | sed "s/$1/token/g" > testdata/marketplace/listings/647415422_curr_abbr=USD_token=token
sleep 1
exit
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/releases/565473?token=$1" | sed "s/$1/token/g" > testdata/releases/565473_token=token
sleep 1
exit
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/releases/1018055?token=$1" | sed "s/$1/token/g" > testdata/releases/1018055_token=token
sleep 1
exit
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/releases/1161277?token=$1" > testdata/releases/1161277_token=token
exit
sleep 1
curl -i -X POST -H "Content-Type:applicaion/json" --user-agent "GoDiscogsTestData" "https://api.discogs.com/marketplace/listings/805055435?token=$1&curr_abr=USD" -d '{"listing_id": 805055435, "release_id":1145342, "condition": "Very Good Plus (VG+)", "price": 5.00, "status":"Draft"}' | sed "s/$1/token/g" >  testdata/marketplace/listings/805055435_curr_abr=USD_token=token
sleep 1
exit
curl -i -X POST -H "Content-Type:applicaion/json" --user-agent "GoDiscogsTestData" "https://api.discogs.com/marketplace/listings/805377159?token=$1&curr_abr=USD" -d '{"listing_id": 805377159, "release_id":11403112, "condition": "Very Good Plus (VG+)", "price": 9.50, "status":"For Sale"}' | sed "s/$1/token/g" >  testdata/marketplace/listings/805377159_curr_abr=USD_token=token
sleep 1
exit
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/marketplace/listings/805377159?curr_abbr=USD&token=$1" | sed "s/$1/token/g" > testdata/marketplace/listings/805377159_curr_abbr=USD_token=token
sleep 1
curl -X DELETE -H "Content-Type:application/json" --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/673768/releases/10866041/instances/280210978?token=$1" | sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/673768/releases/10866041/instances/280210978_token=token
sleep 1
curl -X PUT -H "Content-Type:application/json" --user-agent "GoDiscogsTestData" "https://api.discogs.com/releases/2000000000/rating/BrotherLogic?token=$1" -d '{"rating": 5}'| sed "s/$1/token/g" > testdata/releases/2000000000/rating/brotherlogic_token=token
sleep 1
curl -X PUT -H "Content-Type:application/json" --user-agent "GoDiscogsTestData" "https://api.discogs.com/releases/10567529/rating/BrotherLogic?token=$1" -d '{"rating": 5}'| sed "s/$1/token/g" > testdata/releases/10567529/rating/brotherlogic_token=token
sleep 1
curl -X POST -H "Content-Type:application/json" --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/812802/releases/10?token=$1" -d '' > testdata/users/brotherlogic/collection/folders/812802/releases/10_token=token
sleep 1
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/releases/4707982?token=$1" | sed "s/$1/token/g" > testdata/releases/4707982_token=token
sleep 1
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/releases/11146958?token=$1" | sed "s/$1/token/g" > testdata/users/brotherlogic/collection/releases/11146958_token=token
sleep 1
curl -X POST -H "Content-Type: application/json" --user-agent "GoDiscogsTestData" "https://api.discogs.com//marketplace/listings?token=$1" -d '{"release_id":2576104, "condition":"Very Good Plus (VG+)", "sleeve_condition":"Very Good Plus (VG+)", "price":12.345, "status":"Draft","weight":"auto"}' > testdata/marketplace/listings_token=token
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/marketplace/price_suggestions/2576104?token=$1" |  sed "s/$1/token/g" > testdata/marketplace/price_suggestions/2576104_token=token
sleep 1
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/releases/249504?token=$1" > testdata/releases/249504
sleep 1
curl --user-agent "GoDiscogsTestData" -I "https://api.discogs.com/releases/249504?token=$1" > testdata/releases/249504_token=token.headers
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token
sleep 1
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=2" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=2
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=3" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=3
sleep 1
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=4" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=4
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=5" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=5
sleep 1
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=6" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=6
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=7" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=7
sleep 1
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=8" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=8
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=9" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=9
sleep 1
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=10" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=10
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=11" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=11
sleep 1
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=12" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=12
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=13" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=13
sleep 1
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=14" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=14
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=15" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=15
sleep 1
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=16" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=16
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=17" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=17
sleep 1
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=18" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=18
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=19" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=19
sleep 1
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=20" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=20
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=21" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=21
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=22" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=22
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=23" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=23
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=24" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=24
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=25" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=25
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=26" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=26
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=27" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=27
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=28" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=28
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=29" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=29
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=30" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=30
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=31" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=31
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=32" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=32
sleep 1
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/collection/folders?token=$1" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders_token=token
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/releases/6099374?token=$1" > testdata/releases/6099374_token=token
sleep 1
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/releases/3362530?token=$1" > testdata/releases/3362530_token=token
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/releases/2331612token=$1" > testdata/releases/2331612_token=token
sleep 1
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/releases/228202?token=$1" > testdata/releases/228202_token=token
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/releases/1052359?token=$1" > testdata/releases/1052359_token=token
sleep 1
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/releases/1823781?token=$1" > testdata/releases/1823781_token=token
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/masters/67464/versions?per_page=500&token=$1" > testdata/masters/67464/versions_per_page=500_token=token
sleep 1
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/releases/603365?token=$1" > testdata/releases/603365_token=token
sleep 1
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/releases/9082405?token=$1" > testdata/releases/9082405_token=token
sleep 1
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/releases/2370027?token=$1" > testdata/releases/2370027_token=token
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/masters/38998/versions?per_page=500&token=$1" > testdata/masters/38998/versions_per_page=500_token=token
sleep 1
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/releases/1668957?token=$1" > testdata/releases/1668957_token=token
sleep 1
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/releases/668315?token=$1" > testdata/releases/668315_token=token
sleep 1
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/releases/2425133?token=$1" > testdata/releases/2425133_token=token
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/masters/5251/versions?per_page=500&token=$1" > testdata/masters/5251/versions_per_page=500_token=token
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/wants?per_page=100&token=$1" |  sed "s/$1/token/g" > testdata/users/brotherlogic/wants_per_page=100_token=token
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/wants?per_page=100&token=$1&page=2" |  sed "s/$1/token/g" > testdata/users/brotherlogic/wants_per_page=100_token=token_page=2
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/wants?per_page=100&token=$1&page=3" |  sed "s/$1/token/g" > testdata/users/brotherlogic/wants_per_page=100_token=token_page=3
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/wants?per_page=100&token=$1&page=4" |  sed "s/$1/token/g" > testdata/users/brotherlogic/wants_per_page=100_token=token_page=4
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/wants?per_page=100&token=$1&page=5" |  sed "s/$1/token/g" > testdata/users/brotherlogic/wants_per_page=100_token=token_page=5
sleep 1
curl  --user-agent "GoDiscogsTestData" "https://api.discogs.com/users/brotherlogic/wants?per_page=100&token=$1&page=6" |  sed "s/$1/token/g" > testdata/users/brotherlogic/wants_per_page=100_token=token_page=6

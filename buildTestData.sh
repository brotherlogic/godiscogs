mkdir -p testdata/releases
mkdir -p testdata/users/brotherlogic/collection/folders/0/
mkdir -p testdata/masters/67464
mkdir -p testdata/masters/38998
mkdir -p testdata/masters/5251
curl --user-agent "GoDiscogsTestData" "https://api.discogs.com/releases/249504?token=$1" > testdata/releases/249504
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

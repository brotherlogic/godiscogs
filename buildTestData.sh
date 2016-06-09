mkdir -p testdata/releases
mkdir -p testdata/users/brotherlogic/collection/folders/0/
curl "https://api.discogs.com/releases/249504" > testdata/releases/249504
sleep 1
curl "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token
sleep 1
curl "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=2" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=2
sleep 1
curl "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=3" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=3
sleep 1
curl "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=4" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=4
sleep 1
curl "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=5" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=5
sleep 1
curl "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=6" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=6
sleep 1
curl "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=7" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=7
sleep 1
curl "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=8" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=8
sleep 1
curl "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=9" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=9
sleep 1
curl "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=10" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=10
sleep 1
curl "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=11" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=11
sleep 1
curl "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=12" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=12
sleep 1
curl "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=13" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=13
sleep 1
curl "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=14" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=14
sleep 1
curl "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=15" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=15
sleep 1
curl "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=16" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=16
sleep 1
curl "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=17" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=17
sleep 1
curl "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=18" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=18
sleep 1
curl "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=19" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=19
sleep 1
curl "https://api.discogs.com/users/brotherlogic/collection/folders/0/releases?per_page=100&token=$1&page=20" |  sed "s/$1/token/g" > testdata/users/brotherlogic/collection/folders/0/releases_per_page=100_token=token_page=20

#!/bin/bash

if [ -z "$1" ]
then
  echo "No actual version supplied"
  exit 1
fi

if [ -z "$2" ]
then
  echo "No last version supplied"
  exit 1
fi

if [ -z "$3" ]
then
  echo "No repo github supplied"
  exit 1
fi

sdk_actual_version=$1
sdk_last_version=$2
url=$3
sdk_version=$sdk_actual_version

echo "LAST_VERSION => $sdk_last_version"
echo "ACTUAL_VERSION => $sdk_actual_version"

change_minor=false
while true
do
  #  IFS='.' read -r -a sdk_actual_version_partitionned < "$sdk_actual_version"
  sdk_last_version_partitionned=(`echo $sdk_last_version | tr '.' ' '`)
  major_last="${sdk_last_version_partitionned[0]}"
  minor_last="${sdk_last_version_partitionned[1]}"
  bugfix_last="${sdk_last_version_partitionned[2]}"

  major_last_tag=(`echo $major_last | tr 'v' ' '`)
  echo $major_last_tag
  sdk_version_partitionned=(`echo $sdk_version | tr '.' ' '`)
  #  IFS='.' read -r -a sdk_version_partitionned < "$sdk_version"
  major="${sdk_version_partitionned[0]}"
  minor="${sdk_version_partitionned[1]}"
  bugfix="${sdk_version_partitionned[2]}"

  major_tag=(`echo $major | tr 'v' ' '`)
  echo $major_tag

  echo "toto"

  if [[ $major_tag -gt $major_last_tag ]]
  then
    echo "major tag higher last major tag"
    break
  fi

  if [[ $major_tag -eq $major_last_tag ]] && [[ $minor_actual -gt $minor ]]
  then
      echo "same major. minor is higher"
    break
  fi

  if [[ $major_tag -eq $major_last_tag ]] && [[ $minor_actual -eq $minor ]] && [[ $bugfix_actual -gt $bugfix ]]
  then
    echo "same major/minor. bugfix is higher"
    break
  fi

  if [[ $change_minor = true ]]
  then
    minor_after=$((minor+1))
    sdk_version="$major.$minor_after.0"
    change_minor= false
  else
    bugfix_after=$((bugfix+1))
    sdk_version="$major.$minor.$bugfix_after"
  fi

  return_data=$(curl \
   -H "Accept: application/vnd.github.v3+json" \
   -s https://api.github.com/repos/$url/git/ref/tags/$sdk_version)

  echo "$sdk_version"
  echo $return_data
  is_exist=$(echo $return_data |jq 'has("ref")')
  if [[ $is_exist = true ]]
  then
    git pull
    name_branch="generation-auto-ci-$sdk_version"
    git checkout -B $name_branch
    git config user.email "quentin.cousin76400@gmail.com"
    git config user.name "Repository CI"
    git push --set-upstream origin $name_branch

    ./scripts/update_go_mod.sh $sdk_version
    ./scripts/gen/mock.sh $sdk_version
    echo $sdk_version > ./aws-sdk.version

    git add .
    git commit -m "Automatic ci update Version into ${sdk_version}"
    git push
    git checkout main
    git merge $name_branch
    git tag $sdk_version -m "Update version for ${sdk_version}"
    git push --tags

    echo "new version generated for ${sdk_version}"
    break
  else
    change_minor=true
  fi

done

echo "end of generation missing version"

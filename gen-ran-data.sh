#!/bin/bash

COUNT=$1

# Generate Random Name
#curl -s www.pseudorandom.name

# Generate Random Numbers (to 10)
#for _ in `seq 10`; do echo "$(tr -dc 0-9 </dev/urandom | head -c 10)"; done


#for _ in `seq $COUNT`; do name=$(curl -s www.psuedorandom.name);echo "$name,$(tr -dc 0-9 </dev/urandom | head -c 10)"; done

for _ in `seq $COUNT`; do
  name="$(curl -s www.pseudorandom.name)"
  num="$(tr -dc 0-9 </dev/urandom | head -c 10)"
  echo "$name,$num"
done

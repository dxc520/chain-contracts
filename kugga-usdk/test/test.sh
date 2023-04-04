#!/bin/bash
base_dir=$(cd `dirname $0` && pwd)
cd $base_dir

docker run -it --rm -v `pwd`/../.:/wine-cellar-chain --entrypoint /wine-cellar-chain/test/unit-testing.sh trailofbits/eth-security-toolbox

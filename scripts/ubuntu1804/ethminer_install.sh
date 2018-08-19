#! /bin/bash

# Used to install ethminer from source optimized for nvidia

cd ~ || exit

echo "[INFO] Downloading ethminer"
git clone https://github.com/ethereum-mining/ethminer.git
cd ethmienr || exit

echo "[INFO] Installing dbus development libraries"
sudo apt install libdbus-1-dev
echo "[INFO] Installing cmake"
sudo apt install cmake

echo "[INFO] Updating submodules"
git submodule update --init --recursive
echo "[INFO] Creating build dir"
mkdir build
cd build || exit

echo "[INFO] Configuring cmake with options -DETHASHCUDA=ON -DETHASHCL=OFF -DAPICORE=ON -DETHDBUS=ON -DBINKERN=OFF"
cmake .. -DETHASHCUDA=ON -DETHASHCL=OFF -DAPICORE=ON -DETHDBUS=ON -DBINKERN=OFF

echo "[INFO] Building ethminer project"
cmake --build .

echo "[INFO] Installing executable"
sudo make install
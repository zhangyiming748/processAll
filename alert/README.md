# TTS for Linux
```shell
sudo apt-get install libsndfile1-dev libpulse-dev libncurses5-dev libmp3lame-dev libespeak-dev

wget https://github.com/hgneng/ekho/archive/refs/tags/v9.0.tar.gz
tar xJvf ekho-xxx.tar.xz

// or
git clone https://github.com/hgneng/ekho.git

cd ekho-xxx
./configure --prefix=/usr/local
make
sudo make install
ekho "hello 123"
```

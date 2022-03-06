## Perbaikan yang Dapat Dilakukan (Possible Improvements)
Masalah-masalah yang saya temukan pada program ini adalah:
1. Program tidak bisa mendeteksi gambar yang secara visual sama tetapi memiliki resolusi yang berbeda.
2. Program kurang efisien karena harus membuat hash untuk gambar referensi setiap kali dieksekusi. Hal ini disebabkan oleh tidak adanya database untuk menyimpan hash hasil pemrosesan sebelumnya.
3. Antarmuka yang tidak menarik sehingga pengguna tidak merasa nyaman ketika menggunakan program ini.

Solusi yang dapat diimplementasikan:
1. Memodifikasi metode deteksi gambar yang telah ada. [SKIP]
2. Menggunakan metode lain untuk mendeteksi kesamaan gambar. [DONE: Perceptual Hashing (GoImageHash Lib)]
3. Membuat database yang menyimpan hash gambar referensi.
4. Membuat antarmuka lebih menarik dengan CSS styling. [IN PROGRESS]


# TokoGambar

`TokoGambar` adalah sebuah platform dimana orang bisa menjual gambar yang mereka buat dengan harga tinggi. Gambar itu bisa berupa gambar apapun: wallpaper, fanart, koik, dan lain sebagainya.

Ketika seorang penjual ingin menjual gambarnya, pihak `TokoGambar` akan melakukan beberapa verifikasi untuk menguji bahwa gambar tersebut adalah otentik milik sang penjual. Salah satu langkahnya adalah dengan mengecek apakah gambar tersebut replika dari salah satu gambar yang terdapat di situs `TokoGambar` atau tidak.

Jika gambar tersebut adalah replika, maka gambar tersebut akan ditolak oleh pihak `TokoGambar`.

Suatu gambar dikatakan replika dari gambar lain jika gambar tersebut terlihat sama persis secara visual dengan gambar yang direferensikan. Contoh kasusnya adalah [input_1.jpg](./input/input_1.jpg) & [input_3.jpg](./input/input_3.jpg):

[<img src="./input/input_1.jpg" width="400"/>](./input/input_1.jpg) 

[<img src="./input/input_3.jpg" width="400"/>](./input/input_3.jpg)

Resolusi dari kedua gambar ini berbeda, namun keduanya memiliki visual yang kurang lebih sama. Karena itulah kedua gambar ini dikatakan sebagai replika antara satu sama lain.

Pihak `TokoGambar` menggunakan sistem yang ada di repositori ini untuk mengecek otentisitas dari suatu gambar. Namun sayangnya sistem ini masih berupa prototipe, sehingga masih banyak perbaikan yang perlu dilakukan disana-sini.

## Misi Kamu

Perbaiki sistem yang ada di repositori ini.

Coba pikirkan hal-hal apa saja yang perlu diperbaiki dari sudut pandang bisnis `TokoGambar`. Prioritaskan perubahanmu berdasarkan apa yang memang menurut kamu layak untuk diprioritaskan.

## Pengerjaan Tantangan

1. Fork repositori ini.
2. Analisis masalah dengan menggunakan sudut pandang bisnis `TokoGambar`.
3. Tulis masalah-masalah yang kamu temukan dan apa solusinya di `README.md`.
4. Implementasikan solusi dari masalah yang kamu temukan.
5. Tuliskan juga di `README.md` hal apa saja yang akan kamu lakukan jika kamu punya waktu lebih banyak untuk mengerjakan tantangan ini.
6. Bagikan link repositori hasil fork beserta CV kamu kepada kami.

## Menjalankan Program

Untuk menjalankan program, cukup ketik perintah berikut di terminal:

```
go run *.go
```

Setelah program berjalan, kamu bisa akses programnya di `http://localhost:7124`. Kamu bisa menggunakan gambar-gambar yang ada di direktori [input](./input) untuk mencoba programnya.

Pastikan setidaknya `Go v1.15` sudah terinstall di komputer kamu.

> **Note:**
>
> Implementasi program yang ada di repositori ini masih belum benar. Salah satu indikasinya adalah ketika kamu mencoba gambar `input_1.jpg` & `input_2.jpg` sebagai input di program maka akan keluar hasilnya, sementara kalau kamu coba gambar `input_3.jpg` tidak akan keluar apa-apa.
>
> Kalau implementasi program kamu sudah benar, seharusnya `input_3.jpg` juga akan mengeluarkan hasil yang sama dengan `input_1.jpg` karena keduanya identik secara visual.

## Pertanyaan

Jika ada hal-hal yang ingin ditanyakan, silakan buat issue di repositori ini.

## Pustaka Menarik

- [Perceptual Hash](http://www.hackerfactor.com/blog/index.php?/archives/432-Looks-Like-It.html)
- [Time Complexity Analysis](https://www.youtube.com/watch?v=fZc3ijGM0aM)
- [Levenshtein Distance](https://en.wikipedia.org/wiki/Levenshtein_distance)

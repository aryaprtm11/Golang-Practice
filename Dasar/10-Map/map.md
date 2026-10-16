## Tipe data map

- Pada Array atau Slice, untuk mengakses data, kita menggunakan index Number dimulai dari 0
- Map adalah tipe data lain yang berisikan kumpulan data yang sama, namun kita bisa menentukan jenis tipe data index yang akan kita gunakan
- Sederhananya, Map adalah tipe data kumpulan key-value (kata kunci - nilai), dimana kata kuncinya bersifat unik, tidak boleh sama
- Berbeda dengan Array dan Slice, jumlah data yang kita masukkan ke dalam Map boleh sebanyak-banyaknya, asalkan kata kuncinya berbeda, jika kita gunakan kata kunci sama, maka secara otomatis data sebelumnya akan diganti dengan data baru

## Function Map

| Operasi | Keterangan |
| --- | --- |
| `len(map)` | Mendapatkan jumlah data di map |
| `map[key]` | Mengambil data dari map berdasarkan key |
| `map[key] = value` | Mengubah data di map berdasarkan key |
| `make(map[TypeKey]TypeValue)` | Membuat map baru |
| `delete(map, key)` | Menghapus data di map berdasarkan key |

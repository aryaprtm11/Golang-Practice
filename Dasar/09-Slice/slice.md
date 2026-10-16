## Tipe Data Slice

- Tipe data Slice adalah potongan dari data Array
- Slice mirip dengan Array, yang membedakan adalah ukuran Slice bisa berubah
- Slice dan Array selalu terkoneksi, dimana Slice adalah data yang mengakses sebagian atau seluruh data di Array

## Detail Tipe Data Slice
- Tipe data slice memiliki 3 data, yaitu pointer, length, dan capacity
- Pointer adalah penunjuk data pertama di array pada slice
- Length adalah panjang dari slice
- Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity

## Membuat Slice dari Array

| Membuat Slice | Keterangan |
| --- | --- |
| `array[low:high]` | Membuat slice dari array dimulai indeks `low` sampai indeks sebelum `high` |
| `array[low:]` | Membuat slice dari array dimulai indeks `low` sampai akhir array |
| `array[:high]` | Membuat slice dari array dimulai indeks 0 sampai indeks sebelum `high` |
| `array[:]` | Membuat slice dari array dimulai indeks 0 sampai akhir array |

## Function Slice

| Operasi | Keterangan |
| --- | --- |
| `len(slice)` | Mendapatkan panjang slice |
| `cap(slice)` | Mendapatkan kapasitas slice |
| `append(slice, data)` | Membuat slice baru dengan menambahkan data di posisi terakhir. Jika kapasitas sudah penuh, Go akan membuat array baru. |
| `make([]TypeData, length, capacity)` | Membuat slice baru |
| `copy(destination, source)` | Menyalin data dari `source` ke `destination` |

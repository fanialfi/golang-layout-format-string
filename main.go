package main

import (
	"fmt"
	"golang-layout-format-string/lib"
)

func main() {
	var data = lib.Student{
		Name:        "wick",
		Height:      182.5,
		Age:         26,
		IsGraduated: false,
		Hobies:      []string{"eating", "sleeping"},
	}

	// format speceifer %b untuk memformat nilai integer kedalam bentuk biner (base 2)
	fmt.Printf("umur saya dalam bentuk biner : %b\n", data.Age)

	// format speceifer %c untuk memformat nilai integer ke bentuk string karakter unicode-nya
	// https://www.saidalfaruq.net/artikel/daftar-karakter-unicode-p8z1az6x
	fmt.Printf("it is coffie %c\n", 9749)
	fmt.Printf("%c\n", 9773)

	// format speceifer %d untuk memformat nilai integer menjadi bentuk string numeric base 10
	fmt.Printf("umur saya %d\n", data.Age)

	// format speceifer %f atau %F
	// %F alias untuk %f
	// digunakan untuk memformat data numeric desimal (float) dengan lebar bisa ditentukan
	// secara default lebar digit desimal adalah 6 digit
	fmt.Printf("Height %f\n", data.Height)
	fmt.Printf("Height %F\n", data.Height)

	// format speceifer %o
	// digunakan untuk memformat data numeric kedalam bentuk string numeric berbasis 8 (oktal)
	fmt.Printf("age in oktal %o\n", data.Age)

	// format speceifer %p
	// digunakan untuk memformat data pointer
	// mengembalikan alamat pointer referensi variabel-nya
	// alamat pointer ditulis dalam bentuk hexadesimal dengan prefix 0x
	fmt.Printf("%p\n", &data.Age)

	// format speceifer %T
	// mengembalikan tipe variabel yang akan diformat
	fmt.Printf("%T\n", data.Height)

	// format speceifer %q
	// digunakan untuk ecaping string, meskipun string menggunakan literal \ akan tetap di escape
	// cetak string dengan tanda kutip
	name := "Alice"
	fmt.Printf("Name\t: %s\n", name)
	fmt.Printf("Name\t: %q\n", name)
	name = `fani
	alfirdaus`
	fmt.Printf("%q\n", name)

	// format speceifer %v
	// digunakan untuk memformat data apa saja (termasuk data bertipe any),
	// hasil kembaliannya adalah string dari datanya
	// jika datanya bertipe struct, maka akan ditampilkan semua field-nya secara berurutan
	fmt.Printf("%v\n", data)

	// jika mencetak struct dan ditambahkan plus flags (%+v) maka akan menambahkan nama field-nya
	fmt.Printf("%+v\n", data)

	// format speceifer %x dan %X
	// digunakan untuk memformat data numeric menjadi string numeric berbasis 16 (hexadecimal)
	// speceifer %x dan %X fungsinya sama, perbedaannya %X akan mengembalikan huruf dalam uppercase
	fmt.Printf("age is %x\n", data.Age)
	// jika digunakan pada tipe data string, maka akan menghasilkan kode hexadecimal tiap karakter-nya
	fmt.Printf("name is %X\n", data.Name)

	// format speceifer %%
	// digunakan untuk menampilkan karakter %% pada string
	fmt.Print("%%\n")
}

/* 1)  query untuk menghitung jumlah kata di dalam body pada tabel artikel*/

SELECT sum(array_length(regexp_split_to_array(body, '\s'),1)) FROM artikel;

/* 3)  query untuk menentukan jumlah artikel yang terpublishl*/

SELECT t.* FROM artikel t
LEFT JOIN statuss s ON s.id = t.status 
WHERE s.status = 'publish'

/* 4)  query untuk menampilkan nama author, judul berita, tanggal terbit dan photographer dari 
berita (bila photographer tidak ditemukan diberikan tanda – didalam baris tersebut)
*/
SELECT author.nama AS nama_author, artikel.judul, artikelt.tanggal_terbit, COALESCE(photographer, '-') FROM artikel
LEFT JOIN author ON author.id = artikel.author_id


/*5)  t query untuk menampilkan dalam 1 baris dengan kolom sponsor, judul berita, sumber, 
tanggal terbit dengan format ISO 8601 (bila sponsor tidak ditemukan diberikan tanda – didalam 
baris di kolom tersebut)

*/

SELECT coalesce(meta_artikel.meta_key, '-'), artikel.judul, to_timestamp(tanggal_terbit) AS tanggal_terbits FROM artikel
LEFT JOIN meta_artikel ON meta_artikel.post_id = artikel.id
LIMIT 1


/* 6) rata-rata pageview artikel yang diterbitkan pada bulan januari 2020  */
SELECT AVG(page_view) FROM artikel WHERE tanggal_terbit LIKE '%Januari 2020%'
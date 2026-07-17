package repositories

import (
	"database/sql"
	"one-daily-dose-dispensing-obat-api/structs"
)

func GetLaporanTransaksiObat(db *sql.DB, tanggalAwal string, tanggalAkhir string) ([]structs.LaporanTransaksiObat, error) {

	querySelect := `
		SELECT 	t.id, t.no_resep, p.nama, t.status,
				COUNT(dto.id) AS total_item, t.grand_total AS total_harga, t.tipe_pembayaran,
				t.total_pembayaran, t.kembalian, t.created_at, t.created_by
		FROM transaksi_obat t
		INNER JOIN pasien p ON p.id = t.pasien_id
		LEFT JOIN detail_transaksi_obat dto ON dto.transaksi_obat_id = t.id
		WHERE DATE(t.created_at) BETWEEN $1 AND $2
		GROUP BY 	t.id, t.no_resep, p.nama, t.status,
					t.grand_total, t.tipe_pembayaran, t.total_pembayaran,
					t.kembalian, t.created_at, t.created_by
		ORDER BY t.created_at DESC;
	`

	rows, err := db.Query(querySelect, tanggalAwal, tanggalAkhir)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hasil []structs.LaporanTransaksiObat

	for rows.Next() {

		var setLaporanTransaksiObat structs.LaporanTransaksiObat

		err = rows.Scan(
			&setLaporanTransaksiObat.ID, &setLaporanTransaksiObat.NoResep, &setLaporanTransaksiObat.NamaPasien,
			&setLaporanTransaksiObat.Status, &setLaporanTransaksiObat.TotalItem, &setLaporanTransaksiObat.TotalHarga,
			&setLaporanTransaksiObat.TipePembayaran, &setLaporanTransaksiObat.TotalPembayaran, &setLaporanTransaksiObat.Kembalian,
			&setLaporanTransaksiObat.CreatedAt, &setLaporanTransaksiObat.CreatedBy,
		)

		if err != nil {
			return nil, err
		}

		hasil = append(hasil, setLaporanTransaksiObat)
	}

	if hasil == nil {
		hasil = []structs.LaporanTransaksiObat{}
	}

	return hasil, nil
}

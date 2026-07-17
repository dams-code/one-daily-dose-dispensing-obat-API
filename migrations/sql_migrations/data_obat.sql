
-- +migrate Up
DROP TRIGGER IF EXISTS trigger_users_obat_modified_at ON user_obat;
DROP TRIGGER IF EXISTS trigger_pasien_modified_at ON pasien;
DROP TRIGGER IF EXISTS trigger_obat_modified_at ON obat;
DROP TRIGGER IF EXISTS trigger_kategori_obat_modified_at ON kategori_obat;
DROP TRIGGER IF EXISTS trigger_transaksi_obat_modified_at ON transaksi_obat;
DROP TRIGGER IF EXISTS trigger_generate_kode_obat ON obat;
DROP TRIGGER IF EXISTS trigger_generate_no_resep ON transaksi_obat;
DROP TRIGGER IF EXISTS trigger_generate_no_detail_resep ON detail_transaksi_obat;
DROP TRIGGER IF EXISTS trigger_detail_transaksi_obat_modified_at ON detail_transaksi_obat;

CREATE TABLE IF NOT EXISTS user_obat (
	id SERIAL PRIMARY KEY,
	nama VARCHAR(200) NOT NULL,
	username VARCHAR(50) NOT NULL UNIQUE,
	password VARCHAR(255) NOT NULL,
	role VARCHAR(50) NOT NULL CHECK(role IN ('KEPALAFARMASI', 'ADMINFARMASI','APOTEKER')),
	status BOOLEAN DEFAULT TRUE,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	created_by VARCHAR(50),
	modified_at TIMESTAMP,
	modified_by VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS pasien (
	id SERIAL PRIMARY KEY,
	rm varchar(8) NOT NULL UNIQUE,
	nama VARCHAR(200) NOT NULL,
	jenis_kelamin VARCHAR(200) NOT NULL,
	alamat TEXT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	created_by VARCHAR(50),
	modified_at TIMESTAMP,
	modified_by VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS kategori_obat (
	id SERIAL PRIMARY KEY,
	nama_kategori VARCHAR(200) NOT NULL,
	deskripsi VARCHAR(255),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	created_by VARCHAR(50),
	modified_at TIMESTAMP,
	modified_by VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS obat (
	id SERIAL PRIMARY KEY,
	kategori_id INT REFERENCES kategori_obat(id) ON DELETE SET NULL,
	kode_obat VARCHAR(20) UNIQUE NOT NULL,
	nama_obat VARCHAR(255) NOT NULL,
	stok DECIMAL(10,2) DEFAULT 0,
	satuan VARCHAR(100) NOT NULL,
	harga DECIMAL(10,2) DEFAULT 0,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	created_by VARCHAR(50),
	modified_at TIMESTAMP,
	modified_by VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS transaksi_obat(
	id SERIAL PRIMARY KEY,
	no_resep VARCHAR(50) UNIQUE NOT NULL,
	pasien_id INT NOT NULL REFERENCES pasien(id),	
	status VARCHAR(20) DEFAULT 'PENDING' CHECK(status IN ('PENDING','DISPENSED','CANCELED')),
	tipe_pembayaran VARCHAR(20) CHECK (tipe_pembayaran IN ('TUNAI', 'QRIS')),
	grand_total NUMERIC(18,2) DEFAULT 0 NOT NULL,
    total_pembayaran NUMERIC(18,2),
    kembalian NUMERIC(18,2),
	dibayar_at TIMESTAMP,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	created_by VARCHAR(50),
	modified_at TIMESTAMP,
	modified_by VARCHAR(50)
);

CREATE TABLE detail_transaksi_obat(
    id SERIAL PRIMARY KEY,
    no_detail_resep VARCHAR(20) UNIQUE NOT NULL,
    transaksi_obat_id INT NOT NULL REFERENCES transaksi_obat(id) ON DELETE CASCADE,
    obat_id INT NOT NULL REFERENCES obat(id),
    jumlah DECIMAL(10,2) NOT NULL,
    aturan_pakai VARCHAR(200),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(50),
	modified_at TIMESTAMP,
    modified_by VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS dispensing(
	id SERIAL PRIMARY KEY,
	transaksi_obat_id INT UNIQUE REFERENCES transaksi_obat(id) ON DELETE SET NULL,
	apoteker_id INT REFERENCES user_obat(id) ON DELETE SET NULL,
	tanggal_dispense TIMESTAMP NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION update_pada_modified_at()
RETURNS TRIGGER AS $$
BEGIN
	NEW.modified_at = CURRENT_TIMESTAMP;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +migrate StatementEnd

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION generate_kode_obat()
RETURNS TRIGGER AS
$$
BEGIN
    IF NEW.id IS NULL THEN
        NEW.id := nextval(pg_get_serial_sequence('obat', 'id'));
    END IF;

    NEW.kode_obat := 'OB' || LPAD(NEW.id::TEXT, 4, '0');

    RETURN NEW;
END;
$$
LANGUAGE plpgsql;
-- +migrate StatementEnd

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION generate_no_resep()
RETURNS TRIGGER AS
$$
BEGIN
    IF NEW.id IS NULL THEN
        NEW.id := nextval(pg_get_serial_sequence('transaksi_obat', 'id'));
    END IF;

    NEW.no_resep := 'RSP-' || LPAD(NEW.id::TEXT, 6, '0');

    RETURN NEW;
END;
$$
LANGUAGE plpgsql;
-- +migrate StatementEnd

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION generate_no_detail_resep()
RETURNS TRIGGER AS
$$
BEGIN
    IF NEW.id IS NULL THEN
        NEW.id := nextval(pg_get_serial_sequence('detail_transaksi_obat', 'id'));
    END IF;

    NEW.no_detail_resep := 'DRSP-' || LPAD(NEW.id::TEXT, 6, '0');

    RETURN NEW;
END;
$$
LANGUAGE plpgsql;
-- +migrate StatementEnd

CREATE TRIGGER trigger_users_obat_modified_at
BEFORE UPDATE ON user_obat
FOR EACH ROW
EXECUTE FUNCTION update_pada_modified_at();

CREATE TRIGGER trigger_pasien_modified_at
BEFORE UPDATE ON pasien
FOR EACH ROW
EXECUTE FUNCTION update_pada_modified_at();

CREATE TRIGGER trigger_obat_modified_at
BEFORE UPDATE ON obat
FOR EACH ROW
EXECUTE FUNCTION update_pada_modified_at();

CREATE TRIGGER trigger_kategori_obat_modified_at
BEFORE UPDATE ON kategori_obat
FOR EACH ROW
EXECUTE FUNCTION update_pada_modified_at();

CREATE TRIGGER trigger_transaksi_obat_modified_at
BEFORE UPDATE ON transaksi_obat
FOR EACH ROW
EXECUTE FUNCTION update_pada_modified_at();

CREATE TRIGGER trigger_detail_transaksi_obat_modified_at
BEFORE UPDATE ON detail_transaksi_obat
FOR EACH ROW
EXECUTE FUNCTION update_pada_modified_at();

CREATE TRIGGER trigger_generate_kode_obat
BEFORE INSERT ON obat
FOR EACH ROW
EXECUTE FUNCTION generate_kode_obat();

CREATE TRIGGER trigger_generate_no_resep
BEFORE INSERT ON transaksi_obat
FOR EACH ROW
EXECUTE FUNCTION generate_no_resep();

CREATE TRIGGER trigger_generate_no_detail_resep
BEFORE INSERT ON detail_transaksi_obat
FOR EACH ROW
EXECUTE FUNCTION generate_no_detail_resep();
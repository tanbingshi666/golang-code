package dao

const (
	InsertResourceSQL = `
		INSERT INTO resource (
		id,
		vendor,
		region,
		create_at,
		expire_at,
		type,
		name,
		description,
		status,
		tags,
		update_at,
		sync_at,
		account,
		public_ip,
		private_ip) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`
	InsertResourceDescribeSQL = `
	INSERT INTO resource_describe ( 
	serial_number,
	resource_id, 
	cpu, 
	memory, 
	gpu_amount, 
	gpu_spec, 
	os_type, 
	os_name) VALUES (?, ?, ?, ?, ?, ?, ?, ?);
	`
	QueryOneHostSQL = `
		SELECT
		  t1.id          AS id,
		  t1.vendor      AS vendor,
		  t1.region      AS region,
		  t1.create_at   AS create_at,
		  t1.expire_at   AS expire_at,
		  t1.type        AS type,
		  t1.name        AS name,
		  t1.description AS description,
		  t1.status      AS status,
		  t1.tags        AS tags,
		  t1.update_at   AS update_at,
		  t1.sync_at     AS sync_at,
		  t1.account     AS account,
		  t1.public_ip   AS public_ip,
		  t1.private_ip  AS private_ip,
		  t2.serial_number AS serial_number,
		  t2.resource_id   AS resource_id,
		  t2.cpu           AS cpu,
		  t2.memory        AS memory,
		  t2.gpu_amount    AS gpu_amount,
		  t2.gpu_spec      AS gpu_spec,
		  t2.os_type       AS os_type,
		  t2.os_name       AS os_name
		FROM
		(
		SELECT
		  id,
		  vendor,
		  region,
		  create_at,
		  expire_at,
		  type,
		  name,
		  description,
		  status,
		  tags,
		  update_at,
		  sync_at,
		  account,
		  public_ip,
		  private_ip
		FROM resource WHERE id = ?
		) t1 LEFT JOIN resource_describe AS t2
		ON t1.id = t2.resource_id;
		`

	DeleteOneResourceSQL         = "DELETE FROM resource WHERE id = ?;"
	DeleteOneResourceDescribeSQL = "DELETE FROM resource_describe WHERE serial_number = ?;"

	UpdateOneResourceSQL = "UPDATE resource SET vendor=?, region=?, expire_at=?, name=?, description=? WHERE id = ?;"
)

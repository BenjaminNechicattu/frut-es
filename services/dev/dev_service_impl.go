package devservice

type Health struct {
	IsOk     bool      `json:"is_running"`
	DBStatus *DBStatus `json:"db_running_status"`
}

func (ds *DevService) HealthCheck() (*Health, error) {

	dbStat, err := ds.getDbStatus()
	if err != nil {
		return nil, err
	}

	return &Health{
		IsOk:     true,
		DBStatus: dbStat,
	}, nil
}

type DBStatus struct {
	DBUpTime int `json:"db_up_time"`
}

func (ds *DevService) getDbStatus() (*DBStatus, error) {
	conn, err := ds.mysqlDBService.GetDBInstance()
	if err != nil {
		ds.logger.Error.Print(err)
		return nil, err
	}

	query := `SHOW status LIKE 'Uptime';`

	rows, err := conn.Query(query)
	if err != nil {
		ds.logger.Error.Print(err)
		return nil, err
	}

	var (
		uptime int
		key    string
	)

	for rows.Next() {
		err := rows.Scan(&key, &uptime)
		if err != nil {
			ds.logger.Error.Print(err)
			return nil, err
		}
	}

	return &DBStatus{
		DBUpTime: uptime,
	}, nil
}

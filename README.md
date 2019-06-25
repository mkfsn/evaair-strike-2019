# Eva-air strike 2019

See [Wikipedia](https://zh.wikipedia.org/wiki/2019%E5%B9%B4%E9%95%B7%E6%A6%AE%E8%88%AA%E7%A9%BA%E7%A9%BA%E6%9C%8D%E5%93%A1%E7%BD%B7%E5%B7%A5%E4%BA%8B%E4%BB%B6) for more information.

This is a crawler for parsing time table from [this page](https://booking.evaair.com/flyeva/EVA/B2C/flight-status-erc.aspx?lang=zh-tw&cmstitle=erc-note9&date=20190628-20190705&airport=TPE&Orderby=&reqtime=&ACTCODE=&REASON=) to check if your flight has been canceled or not.

# How to use

```bash
Usage of ./evaair-strike-2019:
	-airport string
		airport to query (default "TPE")
	-end string
		end date to query (default "20190705")
	-start string
		start date to query (default "20190628")
```

# Result

```bash
$ ./evaair-strike-2019  -start 20190701 -end 20190701
{
	"20190701": {
		"Arrival": [
			{
				"Origin": "洛杉磯 (LAX)",
				"Destination": "台北  (TPE)",
				"FlightNumber": "BR015",
				"Status": "取消",
				"Schedule": "2019年6月30日00:50 - 出發2019年7月1日05:45 - 抵達"
			},
			{
				"Origin": "舊金山 (SFO)",
				"Destination": "台北  (TPE)",
				"FlightNumber": "BR027",
				"Status": "取消",
				"Schedule": "2019年6月30日01:20 - 出發2019年7月1日05:30 - 抵達"
			},
			{
				"Origin": "多倫多 (YYZ)",
				"Destination": "台北  (TPE)",
				"FlightNumber": "BR035",
				"Status": "取消",
				"Schedule": "2019年6月30日01:45 - 出發2019年7月1日04:55 - 抵達"
			},
			{
				"Origin": "溫哥華 (YVR)",
				"Destination": "台北  (TPE)",
				"FlightNumber": "BR009",
				"Status": "取消",
				"Schedule": "2019年6月30日02:00 - 出發2019年7月1日05:10 - 抵達"
			},
			{
				"Origin": "巴黎  (CDG)",
				"Destination": "台北  (TPE)",
				"FlightNumber": "BR088",
				"Status": "取消",
				"Schedule": "2019年6月30日11:20 - 出發2019年7月1日06:30 - 抵達"
			},
			{
				"Origin": "洛杉磯 (LAX)",
				"Destination": "台北  (TPE)",
				"FlightNumber": "BR005",
				"Status": "取消",
				"Schedule": "2019年6月30日12:15 - 出發2019年7月1日17:05 - 抵達"
			},
			{
				"Origin": "舊金山 (SFO)",
				"Destination": "台北  (TPE)",
				"FlightNumber": "BR007",
				"Status": "取消",
				"Schedule": "2019年6月30日13:00 - 出發2019年7月1日17:10 - 抵達"
			},
			{
				"Origin": "大阪 (KIX)",
				"Destination": "台北  (TPE)",
				"FlightNumber": "BR179",
				"Status": "取消",
				"Schedule": "2019年6月30日22:10 - 出發2019年7月1日00:10 - 抵達"
			},
			{
				"Origin": "布里斯本 (BNE)",
				"Destination": "台北  (TPE)",
				"FlightNumber": "BR316",
				"Status": "取消",
				"Schedule": "2019年6月30日22:15 - 出發2019年7月1日05:15 - 抵達"
			},
			{
				"Origin": "曼谷 (BKK)",
				"Destination": "台北  (TPE)",
				"FlightNumber": "BR206",
				"Status": "取消",
				"Schedule": "2019年7月1日01:45 - 出發2019年7月1日06:35 - 抵達"
			},
			{
				"Origin": "胡志明市 (SGN)",
				"Destination": "台北  (TPE)",
				"FlightNumber": "BR382",
				"Status": "取消",
				"Schedule": "2019年7月1日01:50 - 出發2019年7月1日06:15 - 抵達"
			},
			{
				"Origin": "馬尼拉 (MNL)",
				"Destination": "台北  (TPE)",
				"FlightNumber": "BR262",
				"Status": "取消",
				"Schedule": "2019年7月1日04:00 - 出發2019年7月1日06:20 - 抵達"
			},
			{
				"Origin": "首爾 (ICN)",
				"Destination": "台北  (TPE)",
				"FlightNumber": "BR149",
				"Status": "取消",
				"Schedule": "2019年7月1日06:50 - 出發2019年7月1日08:25 - 抵達"
			},
			{
				"Origin": "香港 (HKG)",
				"Destination": "台北  (TPE)",
				"FlightNumber": "BR828",
				"Status": "取消",
				"Schedule": "2019年7月1日07:25 - 出發2019年7月1日09:10 - 抵達"
			},
			{
				"Origin": "維也納 (VIE)",
				"Destination": "曼谷 (BKK)",
				"FlightNumber": "BR062",
				"Status": "取消",
				"Schedule": "2019年6月30日18:35 - 出發2019年7月1日10:00 - 抵達"
			},
			{
				"Origin": "曼谷 (BKK)",
				"Destination": "台北  (TPE)",
				"FlightNumber": "BR062",
				"Status": "取消",
				"Schedule": "2019年7月1日11:40 - 出發2019年7月1日16:30 - 抵達"
			},
			{
				"Origin": "倫敦  (LHR)",
				"Destination": "曼谷 (BKK)",
				"FlightNumber": "BR068",
				"Status": "取消",
				"Schedule": "2019年6月30日21:35 - 出發2019年7月1日15:05 - 抵達"
			},
			{
				"Origin": "曼谷 (BKK)",
				"Destination": "台北  (TPE)",
				"FlightNumber": "BR068",
				"Status": "取消",
				"Schedule": "2019年7月1日16:25 - 出發2019年7月1日21:15 - 抵達"
			}
		],
		"Departure": []
	}
}
```

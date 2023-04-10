# Connector

## .env dosya içeriği 
İlk önce ana path' e bir .env dosyası oluşturuyoruz ve içeriğini aşağıdaki gibi dolduruyoruz.
${deger} alanına dbType için gerekli olan parametreyi veriyoruz. 
Her database için farklı lmasına dikkat edelim.

DB_${deger}_HOST
```
DB_MAIN_HOST=127.0.0.1
DB_MAIN_PORT=3306
DB_MAIN_NAME=test
DB_MAIN_USER=mck
DB_MAIN_PASSWORD=123456
```

## main.go dosya içeriği
Şimdilik mysql ve mssql için hazırladım.

```
func main() {
c := connector.Connnector{}

	db, _ := c.CreateConnectionStr("mysql", "MAIN")
	attr, _ := SelectAttributeById(db, 2257)

	println(attr.Name)
}
``
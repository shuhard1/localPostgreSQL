package table

const IndexTmpl = `<!DOCTYPE html>
<html>
<style>
table, th, td {
  border:1px solid black;
}
</style>
<body>
<table style="width:100%">
  <th colspan="2">Order</th>
  <tr>
    <td>order_uid</td>
    <td>{{ .Order_uid }}</td>
  </tr>
  <tr>
    <td>track_number</td>
    <td>{{ .Track_number }}</td>
  </tr>
  <tr>
    <td>entry</td>
    <td>{{ .Entry }}</td>
  </tr>
  <tr>
    <td>delivery</td>
    <td>Name: {{ .Delivery.Name }} <br> 
    Phone: {{ .Delivery.Phone }}<br> 
    Zip: {{ .Delivery.Zip }}<br>
    City: {{ .Delivery.City }}<br>
    Address: {{ .Delivery.Address }}<br>
    Region: {{ .Delivery.Region }}<br>
    Email: {{ .Delivery.Email }}<br></td>
  </tr>
  <tr>
    <td>payment</td>
    <td>Transaction: {{ .Payment.Transaction }} <br> 
    Request_id: {{ .Payment.Request_id }}<br> 
    Currency: {{ .Payment.Currency }}<br>
    Provider: {{ .Payment.Provider }}<br>
    Amount: {{ .Payment.Amount }}<br>
    Payment_dt: {{ .Payment.Payment_dt }}<br>
    Bank: {{ .Payment.Bank }}<br>
    Delivery_cost: {{ .Payment.Delivery_cost }}<br>
    Goods_total: {{ .Payment.Goods_total }}<br>
    Custom_fee: {{ .Payment.Custom_fee }}<br>
	</td>
  </tr>
  <tr>
    <td>items</td>
    <td>{{range .Items}}
    Chrt_id: {{ .Chrt_id }}<br> 
    Track_number: {{ .Track_number }}<br>
    Price: {{ .Price }}<br>
    Rid: {{ .Rid }}<br>
    Name: {{ .Name }}<br>
    Sale: {{ .Sale }}<br>
    Size: {{ .Size }}<br>
    Total_price: {{ .Total_price }}<br>
    Nm_id: {{ .Nm_id }}<br>
    Brand: {{ .Brand }}<br>
    Status: {{ .Status }}<br>
	  {{end}}</td>
  </tr>
  <tr>
    <td>locale</td>
    <td>{{ .Locale }}</td>
  </tr>
  <tr>
    <td>internal_signature</td>
    <td>{{ .Internal_signature }}</td>
  </tr>
  <tr>
    <td>customer_id</td>
    <td>{{ .Customer_id }}</td>
  </tr>
  <tr>
    <td>delivery_service</td>
    <td>{{ .Delivery_service }}</td>
  </tr>
  <tr>
    <td>shardkey</td>
    <td>{{ .Shardkey }}</td>
  </tr>
  <tr>
    <td>sm_id</td>
    <td>{{ .Sm_id }}</td>
  </tr>
  <tr>
    <td>date_created</td>
    <td>{{ .Date_created }}</td>
  </tr>
  <tr>
    <td>oof_shard</td>
    <td>{{ .Oof_shard }}</td>
  </tr>
</table>
</body>
</html>
`

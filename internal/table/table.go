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
    <td>{{range  .Delivery }}</td>
	{{ end }}
  </tr>
  <tr>
    <td>payment</td>
    <td>{{ .Payment }}</td>
  </tr>
  <tr>
    <td>items</td>
    <td>{{ .Items }}</td>
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

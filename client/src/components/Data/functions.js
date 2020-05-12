function formatDate(date) {
  var dd = date.getDate()
  var mm = date.getMonth() + 1
  if (dd < 10) {
    dd = '0' + dd
  }
  if (mm < 10) {
    mm = '0' + mm
  }
  date = dd + '/' + mm
  return date
}

export default formatDate

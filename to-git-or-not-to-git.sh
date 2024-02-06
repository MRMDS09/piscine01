
#!/bin/bash
curl -s https://learn.zone01oujda.ma/assets/superhero/all.json | jq '.[119].name,.[119].powerstats.power ,.[119].appearance.gender'

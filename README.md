# Goban
A small project for my hacky Kanban board implementation in golang

**GET:**   
root-path - / :   
Returns topics of the board

/%topic% :
Returns items of topic

**POST:**   
root-path - / :
Add new topic

/%topic%/{string}
Add new item to topic

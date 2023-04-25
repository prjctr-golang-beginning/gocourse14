## The git flow is not installed in your machine. Just run this command,

macOS:
$ brew install git-flow-avh

Linux:
$ apt-get install git-flow

Windows:
$ wget -q -O - --no-check-certificate https://raw.github.com/petervanderdoes/gitflow-avh/develop/contrib/gitflow-installer.sh install stable | bash

### Ініціалізувати репозиторій з підтримкою Git Flow:
```shell
git flow init
```

### Перейти на гілку розробки:
```shell
git checkout develop
```

### Створити нову фічу (гілку для розробки нового функціоналу):
```shell
git flow feature start feature01
```

### Виконати зміни в коді та комітнути їх:
```shell
git add <file_name>
git commit -m "Feature01: Add new feature"
```
### Запушити зміни до віддаленого репозиторію:
```shell
git push origin feature01
```

### Завершити розробку фічі та змерджити її з гілкою розробки:
```shell
git flow feature finish feature01
```

### Перейти на гілку розробки та змерджити її з гілкою майстра:
```shell
git checkout develop
git merge --no-ff master
```

### Створити реліз (гілку для підготовки до випуску нової версії):
```shell
git flow release start 1.0.0
```

### Виконати зміни в коді для релізу та комітнути їх:
```shell
git add <file_name>
git commit -m "Release 1.0.0: Update version"
```
### Завершити підготовку релізу та змерджити його з гілкою розробки та гілкою майстра:
```shell
git flow release finish 1.0.0
```

### Створити гарячуфікс (гілку для швидкого виправлення помилок в релізі):
```shell
git flow hotfix start hotfix01
```

### Виконати зміни в коді для гарячуфіксу та комітнути їх:
```shell
git add <file_name>
git commit -m "Hotfix01: Fix bug"
```

### Завершити роботу з гарячуфіксом та змерджити його з гілкою розробки та гілкою майстра:
```shell
git flow hotfix finish hotfix01
```

### Виконати синхронізацію з віддаленими гілками:
```shell
git fetch
````
### Оновити локальну копію гілок:
```shell
git pull origin develop
```
`
### Видалити гілку, яка більше не потрібна:
```shell
git branch -d feature01
```
`
### Запушити зміни до віддаленого репозиторію:
```shell
git push
```
`
### Перевірити стан гілок:
```shell
git branch -a
```

### Відправити запит на злиття (Pull Request) для перевірки та злиття з гілкою розробки:

#### Зайти на веб-інтерфейс Git-хостингу та відкрити сторінку зі списком відкритих Pull Request-ів
#### Відкрити Pull Request з відповідною назвою та описом змін
#### Перевірити зміни за допомогою автоматичних перевірок, які виконуються з кожним Pull Request
#### Злити Pull Request з гілкою розробки

### Прибрати локальну копію гілки релізу:
```shell
git branch -d release/1.0.0
```
### Оновити гілку майстра з виправленнями з гарячуфіксу:
```shell
git checkout master
git merge --no-ff hotfix/hotfix01
````
### Запушити зміни до віддаленого репозиторію:
```shell
git push
```
### Видалити гілку гарячуфіксу:
```shell
git branch -d hotfix/hotfix01
```
`
### Оновити гілку розробки з оновленнями з гілки майстра:
```shell
git checkout develop
git merge --no-ff master
````
### Запушити зміни до віддаленого репозиторію:
```shell
git push
```

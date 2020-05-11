sea english text below
# cloudvis
простой консольный клиент для сервиса CloudVision (https://visionbot.ru).
использование: cloudvis опции файл
файл - имя файла изображения в формате png или jpg.
поддерживаемые опции:
-l language
указывает, какой язык использовать, имя языка задаётся в виде двухбуквенного кода, например ru или en. поумолчанию en.
список поддерживаемых языков см. на сайте сервиса.
-stdin
читать изображение со stdin вместо файла.
-t type
указывает тип данных для распознавания. может быть text, image или all. поумолчанию all.
готовые исполняемые файлы для windows вы можете найти в разделе releasies репозитория.
## сборка
требования:
среда разработки [golang](https://golang.org)
тестировалось на go 1.14.1
дополнительные пакеты не требуются

затем
go get github.com/ilyapashuk/cloudvis

процесс распознавания осуществляется в облаке при использовании [этого](https://visionbot.ru) сервиса.
это приложение не осуществляет какую-либо обработку загружаемых изображений, поэтому обо всех ошибках, связанных с некорректным распознаванием, следует сообщать напрямую [разработчику](mailto:aleks-samos@yandex.ru) сервиса.
обратите также внимание, что хотя распознавание текста работает достаточно хорошо, не стоит полностью доверять результату в вопросе распознавания предметов.

в своей работе приложение использует следующий api:
https://visionbot.ru/apiv2

## english
simple console client for CloudVision service (https://visionbot.ru).
Usage: cloudvis opts file 
 file is a image file name. png and jpg are officially supported. 
 Opts can be one of:
  -l language
    	set language name to use. language is formed as two letter code, like ru or en, default is en. (default "en")
  -stdin
    	get image data from stdin, filename will be ignored
  -t type
    	data type to recognize. can be 'text', 'image' or 'all'. (default "all")
precompiled binaries for windows can be found in releases section
## build
requirements:
[golang](https://golang.org) build environment.
tested on go 1.14.1
no any additional packages required.

then
go get github.com/ilyapashuk/cloudvis

all recognition computations are made in the cloud by [this](https://visionbot.ru) service.
this app don't makes any preprocessing for images, so all wrong recognition issues should be passed to [service's [developer](mailto:aleks-samos@yandex.ru)

in it's work this app uses following api:
https://visionbot.ru/apiv2

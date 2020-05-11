sea english text below
# cloudvis
������� ���������� ������ ��� ������� CloudVision (https://visionbot.ru).
�������������: cloudvis ����� ����
���� - ��� ����� ����������� � ������� png ��� jpg.
�������������� �����:
-l language
���������, ����� ���� ������������, ��� ����� ������� � ���� �������������� ����, �������� ru ��� en. ����������� en.
������ �������������� ������ ��. �� ����� �������.
-stdin
������ ����������� �� stdin ������ �����.
-t type
��������� ��� ������ ��� �������������. ����� ���� text, image ��� all. ����������� all.
������� ����������� ����� ��� windows �� ������ ����� � ������� releasies �����������.
## ������
����������:
����� ���������� [golang](https://golang.org)
������������� �� go 1.14.1
�������������� ������ �� ���������

�����
go get github.com/ilyapashuk/cloudvis

������� ������������� �������������� � ������ ��� ������������� [�����](https://visionbot.ru) �������.
��� ���������� �� ������������ �����-���� ��������� ����������� �����������, ������� ��� ���� �������, ��������� � ������������ ��������������, ������� �������� �������� [������������](mailto:aleks-samos@yandex.ru) �������.
�������� ����� ��������, ��� ���� ������������� ������ �������� ���������� ������, �� ����� ��������� �������� ���������� � ������� ������������� ���������.

� ����� ������ ���������� ���������� ��������� api:
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

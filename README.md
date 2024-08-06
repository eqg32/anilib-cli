# anilib-cli

## Что?
Это клиент сайта anilib. Я абсолютно никак не связан с его администрацией. Чтобы получить мелкую спраку по использованию этой программы просто введите:
```bash
./anilib-cli -h
```
Эта программа не проигрывает видео из Kodik. Я вряд ли буду писать его поддержку, ибо мне лень.

## Компиляция
Для компиляции нужен лишь установленный Go.
```bash
go build .
```
В рабочей папке появится бинарник anilib-cli.
**Вы великолепны!**

## Использование
В конечном итоге эта программка выдает просто ссылку на видео. Чтоб это видео проиграть, нужен плеер. Сойдет `mpv`, `vlc` и тд.
Вот два примера включения первой серии Евангелиона:
```bash
# mpv
./anilib-cli --search Evangelion --anime 1 --episode 1 --video 4 | mpv --playlist=-
# vlc
vlc $(sh -c "./anilib-cli --search Evangelion --anime 1 --episode 1 --video 4")
```

Теперь можно дописать флаг `--mpv` вместо `| mpv --playlist=-`:
```bash
./anilib-cli --search Evangelion --anime 1 --episode 1 --video 4 --mpv
```

Если хотите просто поискать по названию, то команда будет выглядеть как-то так:
```bash
./anilib-cli --search "Wolf's Rain"
```
Аргументов всего 5, потому разобраться в программе достаточно просто.

Приятного использования штолеее.
Смайлики, чтоб этот `README.md` выглядил айтишно: 😃😁😀😅😊😇😂🤗🫢🤫🤭🤙🤙🤙

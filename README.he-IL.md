<h1 align="center">בַנגו</h1>

<p align="center">חיפוש עם בנגים</p>

<p align="center">
    [<a href="https://github.com/TheAnnoying/bango/blob/main/README.md">English</a>]
</p>

> הבנגים נלקחו מ[הליום](https://helium.computer/bangs).

> [!TIP]
> אחרי הרצה, הוסיפו את בנגו לדפדפן שלכם בעזרת קישור במבנה הזה:
>
> ```
> http://localhost:8479/?q=%s
> ```
> אם אתם לא רוצים להריץ את בנגו בעצמכם, אפשר להשתמש בגרסה הציבורית שלי:
>
> ```
> https://bango.theannoying.dev/?q=%s
> ```

### משתני סביבה
| משתנה | שימוש | ערך ברירת מחדל |
| -------- | ----- | ------------- |
| PORT | הפורט שעליו בנגו ירוץ | 8479 |
| BANG_PREFIX | הקידומת של בנגים בשאילתות חיפוש | ! |
| DEFAULT_SEARCH_PREFIX | החלק של הקישור שמגיע לפני שאילתת החיפוש | `https://www.google.com/search?q=` |
| DEFAULT_SEARCH_SUFFIX | שאר הקישור, החלק שאחרי שאילתת החיפוש | כלום. דוגמאות: `&submit_search=`, `&searchon=all&suite=all&section=all` |

### דוקר
צרו קובץ הנקרא `docker-compose.yml` עם התוכן הבא:
```yml
services:
  bango:
    image: ghcr.io/theannoying/bango:latest
    ports:
      - "8479:8479"
    restart: unless-stopped
```
הריצו בעזרת הפקודה: `docker compose up -d`
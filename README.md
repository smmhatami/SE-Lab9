# گزارش آزمایش نهم درس مهندسی نرم افزار

## اعضای گروه

+ سید محمدمهدی حاتمی 98109561
+ پیمان حاجی محمد 98170776

## شرح آزمایش 

در گام اول یک RESTful API را پیاده سازی می کنیم که یک عملیات CRUD ساده را ارائه می‌دهد. برای پیاده سازی این قسمت از زبان Go و کتابخانه ی Gin استفاده می کنیم. در این پروژه یک موجودیت User داریم که حاوی نام و نام خانوادگی است و عملیات های CRUD را برای این موجودیت پیاده سازی کرده و API آن را تعریف کرده ایم. در ضمن از یک پایگاه داده ی Postgres SQL به منظور ذخیره ی داده ی کاربران استفاده می کنیم. 

![Screenshot from 2024-01-12 22-58-06](https://github.com/smmhatami/SE-Lab9/assets/62210297/66273cc0-c37a-46ff-b0b4-9a0537576be7)

سپس به سراغ dockerise نمودن این پروژه می رویم. داکرفایل تعریف شده برای این پروژه نکته ی خاصی ندارد. در docker-compose هم برای پایگاه داده از سرور های آماده استفاده کرده و برای backend پروژه ی خودمان را build  می کنیم. 

![Screenshot from 2024-01-12 23-00-37](https://github.com/smmhatami/SE-Lab9/assets/62210297/1acfc611-2954-4d94-b9c6-9162c07c6478)

در گام بعدی برای اجرای همزمان چند نسخه از سرور و تقسیم بار بین آن ها، یک سرویس nginx نیز به پروژه اضافه می کنیم. این سرویس را نیز در docker-compose اضافه کرده و یک network برای اتصال این موجودیت ها به یکدیگر نیز تعریف می کنیم. حال با اضافه کرده فیلد replicas به سرویس backend در فایل docker-compose کاری می کنیم که چند نسخه از سرور به صورت همزمان اجرا شوند. در تصویر زیر نتیجه ی اجرای docker-compose را مشاهده می کنید. 

![Screenshot from 2024-01-12 23-05-03](https://github.com/smmhatami/SE-Lab9/assets/62210297/e7be0fb4-33ed-402e-82a2-9541cf1a470d)

همینطور که مشاهده می کنید چند نسخه ی مختلف از سرور ما در حال اجرا هستند. همچنین با اجرای دستورات docker container ls و docker image ls میتوانیم کانتینر های در حال اجرا و image های بیلد شده را مشاهده کنیم. 

![Screenshot from 2024-01-12 23-06-55](https://github.com/smmhatami/SE-Lab9/assets/62210297/548fa101-24ba-48e4-8747-8f54c056526d)

مطابق تصویر مشاهده می کنیم که تنها یک image از crud-backend داریم ولی ۴ نسخه از آن در حال اجرا هستند. 

همچنین اگر با استفاده postman درخواستی برای سرور ارسال کنیم، مشاهده می کنیم که توزیع بار انجام شده و درخواست ها توسط کارگزار های مختلفی پاسخ داده می شود، ولی پایگاه داده ی همه ی آن ها یکی است. 

نتیجه ی چند بار فراخوانی گرفتن تمام کاربران : 

![Screenshot from 2024-01-12 23-10-15](https://github.com/smmhatami/SE-Lab9/assets/62210297/2e7bb610-a0b0-4c2a-aea0-56fd48e25402)

![Screenshot from 2024-01-12 23-10-38](https://github.com/smmhatami/SE-Lab9/assets/62210297/0e89a09a-1537-4bcb-91cd-5b1567acbe48)


نتیجه ی چند بار فراخوانی ساخت کاربر جدید :‌

![Screenshot from 2024-01-12 23-11-52](https://github.com/smmhatami/SE-Lab9/assets/62210297/2cd59158-05a6-4fe6-890b-2f3b36a665c3)

![Screenshot from 2024-01-12 23-12-11](https://github.com/smmhatami/SE-Lab9/assets/62210297/1d11b86c-9f0d-4f76-a64b-29ff73389028)

فراخوانی مجدد گرفتن تمام کاربران : 

![Screenshot from 2024-01-12 23-13-08](https://github.com/smmhatami/SE-Lab9/assets/62210297/0a897443-b023-4ad2-9a8c-b4c75c62a779)

![Screenshot from 2024-01-12 23-13-17](https://github.com/smmhatami/SE-Lab9/assets/62210297/f718192f-bc83-44de-a9b5-d16ebee1a9e4)

با بالا رفتن فشار وارده به backend پروژه ی ما کافیست که با توجه به میزان فشار تعداد نمونه های در حال اجرای سرور را افزایش دهیم. به این منظور در فایل docker-compose تعداد replicas را تغییر می دهیم. 

![Screenshot from 2024-01-12 23-18-30](https://github.com/smmhatami/SE-Lab9/assets/62210297/68a4d77f-a4e1-4a3b-9270-bed97b520801)

## پرسش ها 
در محیط‌ های Docker و معماری میکروسرویس، مفهوم stateless به معنای عدم وجود حالت مشخص در یک سیستم یا سرویس است. این مفهوم به این معناست که هر واحد سرویس یا کانتینر Docker مستقل از سایر واحدها و بدون وابستگی به حالت یا داده‌ های قبلی اجرا می‌ شود.
در Docker، اگر یک کانتینر stateless باشد، به این معناست که هیچ تغییر دائمی در فایل‌ ها یا داده‌ های داخلی آن ایجاد نمی‌ شود. همه تغییرات و اطلاعات مورد نیاز برای اجرا در زمان اجرا (runtime) انجام می‌ شود. این ویژگی باعث می‌شود که کانتینرها به راحتی قابل مقیاس‌ پذیری باشند و در محیط‌ های متنوع اجرا شوند.
در این آزمایش نیز ما سرویس های بکند خود را بعلت stateless بودن آنها توانستیم به صورت مستقل اجرا کرده و توسط load balancer آنها را مدیریت کنیم.

package target

import "strategy/constant"

var 
	I18Map  = map[int]I18{}
	I18List = []I18{
		{
			Num: 1,
			Img: ``,
			ZH:  `用户名`,
			En:  `User Name`,
			Ru:  `Имя пользователя`,
			DE:  `Benutzername`,
			FR:  `Nom d'utilisateur`,
		}, {
			Num: 2,
			Img: ``,
			ZH:  `请输入用户名`,
			En:  `Please input user name`,
			Ru:  `Введите имя пользователя`,
			DE:  `Bitte Benutzernamen eingeben`,
			FR:  `Veuillez renseigner votre nom d'utilisateur`,
		}, {
			Num: 3,
			Img: ``,
			ZH:  `密码`,
			En:  `Password`,
			Ru:  `Пароль`,
			DE:  `Passwort`,
			FR:  `Mot de passe`,
		}, {
			Num: 4,
			Img: ``,
			ZH:  `请输入登录密码`,
			En:  `Please input password for login`,
			Ru:  `Введите пароль для входа`,
			DE:  `Bitte gib das Passwort zum Anmelden ein`,
			FR:  `Veuillez renseigner votre mot de passe afin de vous connecter`,
		}, {
			Num: 5,
			Img: ``,
			ZH:  `请输入邮箱验证码`,
			En:  `Please input Email verification code`,
			Ru:  `Введите код подтверждения адреса электронной почты`,
			DE:  `Bitte gib den Email Verifizierungscode ein`,
			FR:  `Veuillez renseigner le code de vérification reçu par Email`,
		}, {
			Num: 6,
			Img: ``,
			ZH:  `发送验证码`,
			En:  `Send Verification Code`,
			Ru:  `Отправить код подтверждения`,
			DE:  `Verifizierungscode senden`,
			FR:  `Envoyer le code de vérification`,
		}, {
			Num: 7,
			Img: ``,
			ZH:  `请前往游戏内邮箱获取验证码`,
			En:  `Please check the verification code sent to the in-game mailbox`,
			Ru:  `Проверьте код подтверждения, отправленный на внутриигровой почтовый ящик`,
			DE:  `Bitte überprüft den Verifizierungscode, der an das Postfach im Spiel gesendet wird`,
			FR:  `Nous avons envoyé un code de vérification dans votre boîte de réception en jeu`,
		}, {
			Num: 8,
			Img: ``,
			ZH:  `登录`,
			En:  `Login`,
			Ru:  `Войти`,
			DE:  `Anmelden`,
			FR:  `Se connecter`,
		}, {
			Num: 9,
			Img: ``,
			ZH:  `注册账户`,
			En:  `Register`,
			Ru:  `Регистрация`,
			DE:  `Registrierung`,
			FR:  `S'inscrire`,
		}, {
			Num: 10,
			Img: ``,
			ZH:  `忘记密码`,
			En:  `Forget Password?`,
			Ru:  `Забыли пароль?`,
			DE:  `Passwort vergessen?`,
			FR:  `Mot de passe oublié ?`,
		}, {
			Num: 11,
			Img: ``,
			ZH:  `登录`,
			En:  `Login`,
			Ru:  `Войти`,
			DE:  `Anmelden`,
			FR:  `Se connecter`,
		}, {
			Num: 12,
			Img: ``,
			ZH:  `注册`,
			En:  `Register`,
			Ru:  `Регистрация`,
			DE:  `Registrierung`,
			FR:  `S'inscrire`,
		}, {
			Num: 13,
			Img: ``,
			ZH:  `用户名`,
			En:  `User Name`,
			Ru:  `Имя пользователя`,
			DE:  `Benutzername`,
			FR:  `Nom d'utilisateur`,
		}, {
			Num: 14,
			Img: ``,
			ZH:  `请输入纯字母或数字 不超过10个字符`,
			En:  `Please input letters or numbers only with a maximum of 10 characters.`,
			Ru:  `Вводите только буквы или цифры длиной не более 10 символов.`,
			DE:  `Bitte gebt nur Buchstaben oder Zahlen mit maximal 10 Zeichen ein.`,
			FR:  `Veuillez utiliser uniquement des lettres ou des chiffres avec une limite de 10 caractères maximum.`,
		}, {
			Num: 15,
			Img: ``,
			ZH:  `游戏ID`,
			En:  `Game ID`,
			Ru:  `Игровой ID`,
			DE:  `Spiel ID`,
			FR:  `ID du joueur`,
		}, {
			Num: 16,
			Img: ``,
			ZH:  `请输入游戏ID`,
			En:  `Please input your game ID`,
			Ru:  `Введите ваш игровой ID`,
			DE:  `Bitte gib deine Spiel ID ein`,
			FR:  `Veuillez renseigner votre ID de joueur`,
		}, {
			Num: 17,
			Img: ``,
			ZH:  `区服ID`,
			En:  `Server ID`,
			Ru:  `ID сервера`,
			DE:  `Server ID`,
			FR:  `ID du serveur`,
		}, {
			Num: 18,
			Img: ``,
			ZH:  `请输入区服ID`,
			En:  `Please input your server ID`,
			Ru:  `Введите ID вашего сервера`,
			DE:  `Bitte gib deine Server ID ein`,
			FR:  `Veuillez renseigner votre ID de serveur`,
		}, {
			Num: 19,
			Img: ``,
			ZH:  `请输入邮箱验证码`,
			En:  `Please input Email verification code`,
			Ru:  `Введите код подтверждения адреса электронной почты`,
			DE:  `Bitte gib den Email Verifizierungscode ein`,
			FR:  `Veuillez renseigner le code de vérification reçu par Email`,
		}, {
			Num: 20,
			Img: ``,
			ZH:  `发送验证码`,
			En:  `Send Verification Code`,
			Ru:  `Отправить код подтверждения`,
			DE:  `Verifizierungscode senden`,
			FR:  `Envoyer le code de vérification`,
		}, {
			Num: 21,
			Img: ``,
			ZH:  `请前往游戏内邮箱获取验证码`,
			En:  `Please check the verification code sent to the in-game mailbox`,
			Ru:  `Проверьте код подтверждения, отправленный на внутриигровой почтовый ящик`,
			DE:  `Bitte überprüft den Verifizierungscode, der an das Postfach im Spiel gesendet wird`,
			FR:  `Nous avons envoyé un code de vérification dans votre boîte de réception en jeu`,
		}, {
			Num: 22,
			Img: ``,
			ZH:  `密码`,
			En:  `Password`,
			Ru:  `Пароль`,
			DE:  `Passwort`,
			FR:  `Mot de passe`,
		}, {
			Num: 23,
			Img: ``,
			ZH:  `请输入登录密码`,
			En:  `Please input password for login`,
			Ru:  `Введите пароль для входа`,
			DE:  `Bitte gib das Passwort zum Anmelden ein`,
			FR:  `Veuillez renseigner votre mot de passe afin de vous connecter`,
		}, {
			Num: 24,
			Img: ``,
			ZH:  `密码确认`,
			En:  `Confirm Password`,
			Ru:  `Подтвердите пароль`,
			DE:  `Passwort bestätigen`,
			FR:  `Confirmer le mot de passe`,
		}, {
			Num: 25,
			Img: ``,
			ZH:  `请再次输入登录密码`,
			En:  `Please input the password again`,
			Ru:  `Пожалуйста, введите пароль еще раз`,
			DE:  `Bitte gib das Passwort erneut ein`,
			FR:  `Veuillez renseigner votre mot de passe à nouveau`,
		}, {
			Num: 26,
			Img: ``,
			ZH:  `注册并登录`,
			En:  `Register and Login`,
			Ru:  `Зарегистрируйтесь и войдите в систему`,
			DE:  `Registrierung und Anmeldung`,
			FR:  `S'inscrire et se connecter`,
		}, {
			Num: 27,
			Img: ``,
			ZH:  `修改密码`,
			En:  `Change Password`,
			Ru:  `Изменить пароль`,
			DE:  `Passwort ändern`,
			FR:  `Changer le mot de passe`,
		}, {
			Num: 28,
			Img: ``,
			ZH:  `新密码`,
			En:  `New Password`,
			Ru:  `Новый пароль`,
			DE:  `Neues Passwort`,
			FR:  `Nouveau mot de passe`,
		}, {
			Num: 29,
			Img: ``,
			ZH:  `请输入新的登录密码`,
			En:  `Please input the new password`,
			Ru:  `Введите новый пароль`,
			DE:  `Bitte gib das neue Passwort ein`,
			FR:  `Veuillez renseigner votre nouveau mot de passe`,
		}, {
			Num: 30,
			Img: ``,
			ZH:  `密码确认`,
			En:  `Confirm Password`,
			Ru:  `Подтвердите пароль`,
			DE:  `Passwort bestätigen`,
			FR:  `Confirmer le mot de passe`,
		}, {
			Num: 31,
			Img: ``,
			ZH:  `游戏ID`,
			En:  `Game ID`,
			Ru:  `Игровой ID`,
			DE:  `Spiel ID`,
			FR:  `ID du joueur`,
		}, {
			Num: 32,
			Img: ``,
			ZH:  `请输入游戏ID`,
			En:  `Please input your game ID`,
			Ru:  `Введите ваш игровой ID`,
			DE:  `Bitte gib deine Spiel ID ein`,
			FR:  `Veuillez renseigner votre ID de joueur`,
		}, {
			Num: 33,
			Img: ``,
			ZH:  `区服ID`,
			En:  `Server ID`,
			Ru:  `ID сервера`,
			DE:  `Server ID`,
			FR:  `ID du serveur`,
		}, {
			Num: 34,
			Img: ``,
			ZH:  `请输入区服ID`,
			En:  `Please input your server ID`,
			Ru:  `Введите ID вашего сервера`,
			DE:  `Bitte gib deine Server ID ein`,
			FR:  `Veuillez renseigner votre ID de serveur`,
		}, {
			Num: 35,
			Img: ``,
			ZH:  `请输入邮箱验证码`,
			En:  `Please input Email verification code`,
			Ru:  `Введите код подтверждения адреса электронной почты.`,
			DE:  `Bitte gib den Email Verifizierungscode ein`,
			FR:  `Veuillez renseigner le code de vérification reçu par Email`,
		}, {
			Num: 36,
			Img: ``,
			ZH:  `发送验证码`,
			En:  `Send verification Code`,
			Ru:  `Отправить код подтверждения`,
			DE:  `Verifizierungscode senden`,
			FR:  `Envoyer le code de vérification`,
		}, {
			Num: 37,
			Img: ``,
			ZH:  `请前往游戏内邮箱获取验证码`,
			En:  `Please check the verification code sent to the in-game mailbox`,
			Ru:  `Проверьте код подтверждения, отправленный на внутриигровой почтовый ящик`,
			DE:  `Bitte überprüft den Verifizierungscode, der an das Postfach im Spiel gesendet wird`,
			FR:  `Nous avons envoyé un code de vérification dans votre boîte de réception en jeu`,
		}, {
			Num: 38,
			Img: ``,
			ZH:  `修改密码`,
			En:  `Change Password`,
			Ru:  `Изменить пароль`,
			DE:  `Passwort ändern`,
			FR:  `Changer le mot de passe`,
		}, {
			Num: 39,
			Img: ``,
			ZH:  `提示`,
			En:  `Notification`,
			Ru:  `Уведомление`,
			DE:  `Benachrichtigung`,
			FR:  `Notification`,
		}, {
			Num: 40,
			Img: ``,
			ZH:  `当前登录账户因违反社区规则`,
			En:  `Due to the violation of the community rules by the current account`,
			Ru:  `В связи с нарушением правил сообщества текущим аккаунтом`,
			DE:  `Aufgrund der Verletzung der Gemeinschaftsregelung durch das aktuelle Konto`,
			FR:  `En raison de la violation des règles de la communauté par le compte actuel`,
		}, {
			Num: 41,
			Img: ``,
			ZH:  `账户已经封停至xxx年xx月xx日xx:xx`,
			En:  `This account has been suspended untill XX(month)/XX(day)/XXXX(year) XX:XX`,
			Ru:  `Эта учетная запись была приостановлена до XX(месяц)/XX(день)/XX(год) XX:XX`,
			DE: `Dieses Konto wurde bis zum XX(Monat)/XX(Tag)/XXXX(Jahr) XX:XX gesperrt.
				DE: `Dieses Konto wurde bis zum XX(Monat)/XX(Tag)/XXXX(Jahr) XX:XX gesperrt.
			FR: `Ce compte a été suspendu jusqu'au XX(mois)/XX(jour)/XXXX(année) XX:XX (In France as well, we use a DAY/MONTH/YEAR format)`,
		}, {
			Num: 42,
			Img: ``,
			ZH: `密码修改成功
				ZH: `密码修改成功
			En: `The password has been successfully changed. Please use the new password to login.`,
			Ru: `Пароль успешно изменен. Пожалуйста, используйте новый пароль для входа.`,
			DE: `Das Passwort wurde erfolgreich geändert. Bitte verwende das neue Passwort zum Anmelden.`,
			FR: `Le mot de passe a été modifié avec succès. Veuillez utiliser le nouveau mot de passe pour vous connecter.`,
		}, {
			Num: 43,
			Img: ``,
			ZH:  `登录`,
			En:  `Login`,
			Ru:  `Войти`,
			DE:  `Anmelden`,
			FR:  `Se connecter`,
		}, {
			Num: 44,
			Img: ``,
			ZH:  `请先登录 再进行操作`,
			En:  `Please login first before taking other actions`,
			Ru:  `Пожалуйста, войдите в систему, прежде чем предпринимать другие действия`,
			DE:  `Bitte melde dich zuerst an, bevor du weitere Aktionen durchführst`,
			FR:  `Veuillez d'abord vous connecter avant de réaliser cette action`,
		}, {
			Num: 45,
			Img: ``,
			ZH:  `用户名`,
			En:  `User Name`,
			Ru:  `Имя пользователя`,
			DE:  `Benutzername`,
			FR:  `Nom d'utilisateur`,
		}, {
			Num: 46,
			Img: ``,
			ZH:  `请输入用户名`,
			En:  `Please input user name`,
			Ru:  `Введите имя пользователя`,
			DE:  `Bitte Benutzernamen eingeben`,
			FR:  `Veuillez renseigner votre nom d'utilisateur`,
		}, {
			Num: 47,
			Img: ``,
			ZH:  `密码`,
			En:  `Password`,
			Ru:  `Пароль`,
			DE:  `Passwort`,
			FR:  `Mot de passe`,
		}, {
			Num: 48,
			Img: ``,
			ZH:  `请输入登录密码`,
			En:  `Please input password for login`,
			Ru:  `Введите пароль для входа`,
			DE:  `Bitte gib das Passwort zum Anmelden ein`,
			FR:  `Veuillez renseigner votre mot de passe afin de vous connecter`,
		}, {
			Num: 49,
			Img: ``,
			ZH:  `请输入邮箱验证码`,
			En:  `Please input Email verification code`,
			Ru:  `Введите код подтверждения адреса электронной почты.`,
			DE:  `Bitte gib den Email Verifizierungscode ein`,
			FR:  `Veuillez renseigner le code de vérification reçu par Email`,
		}, {
			Num: 50,
			Img: ``,
			ZH:  `发送验证码`,
			En:  `Send Verification Code`,
			Ru:  `Отправить код подтверждения`,
			DE:  `Verifizierungscode senden`,
			FR:  `Envoyer le code de vérification`,
		}, {
			Num: 51,
			Img: ``,
			ZH:  `请前往游戏内邮箱获取验证码`,
			En:  `Please check the verification code sent to the in-game mailbox`,
			Ru:  `Проверьте код подтверждения, отправленный на внутриигровой почтовый ящик`,
			DE:  `Bitte überprüft den Verifizierungscode, der an das Postfach im Spiel gesendet wird`,
			FR:  `Nous avons envoyé un code de vérification dans votre boîte de réception en jeu`,
		}, {
			Num: 52,
			Img: ``,
			ZH:  `注册账户`,
			En:  `Register Account`,
			Ru:  `Зарегистрировать аккаунт`,
			DE:  `Konto registrieren`,
			FR:  `Créer un compte`,
		}, {
			Num: 53,
			Img: ``,
			ZH:  `攻略制作器`,
			En:  `Guide Generator`,
			Ru:  `Создание руководств`,
			DE:  `Leitfaden-Generator`,
			FR:  `Générateur de guide`,
		}, {
			Num: 54,
			Img: ``,
			ZH:  `我的书签`,
			En:  `My Bookmarks`,
			Ru:  `Мои закладки`,
			DE:  `Meine Lesezeichen`,
			FR:  `Mes favoris`,
		}, {
			Num: 55,
			Img: ``,
			ZH:  `搜索`,
			En:  `Search`,
			Ru:  `Поиск`,
			DE:  `Suche`,
			FR:  `Rechercher`,
		}, {
			Num: 56,
			Img: ``,
			ZH:  `筛选`,
			En:  `Reset`,
			Ru:  `Перезагрузить`,
			DE:  `Zurücksetzen`,
			FR:  `Réinitialiser`,
		}, {
			Num: 57,
			Img: ``,
			ZH:  `全部重置`,
			En:  `Reset All`,
			Ru:  `Сбросить все`,
			DE:  `Alles zurücksetzen`,
			FR:  `Tout réinitialiser`,
		}, {
			Num: 58,
			Img: ``,
			ZH:  `选择游戏关卡`,
			En:  `Please Select  Game Stage`,
			Ru:  `Пожалуйста, выберите стадию игры`,
			DE:  `Bitte Spielphase wählen`,
			FR:  `Veuillez sélectionner l'étape de jeu`,
		}, {
			Num: 59,
			Img: ``,
			ZH:  `显示所有早期游戏内容 推荐攻略内容`,
			En:  `Recommended Guide Contents`,
			Ru:  `Рекомендуемое содержание руководства`,
			DE:  `Empfohlene Inhalte des Leitfadens`,
			FR:  `Contenus de guide recommandés`,
		}, {
			Num: 60,
			Img: ``,
			ZH:  `最新`,
			En:  `Newest`,
			Ru:  `Новейший`,
			DE:  `Neuestes`,
			FR:  `Récent`,
		}, {
			Num: 61,
			Img: ``,
			ZH:  `最热`,
			En:  `Most Popular`,
			Ru:  `Самый популярный`,
			DE:  `Am beliebtesten`,
			FR:  `Plus populaire`,
		}, {
			Num: 62,
			Img: ``,
			ZH:  `点赞`,
			En:  `Likes`,
			Ru:  `Мне нравится`,
			DE:  `(gefallen) Like ("gefallen" is correct if you look at the language but as a Thumbsup button you would still use the english "like")`,
			FR:  `J'aime`,
		}, {
			Num: 63,
			Img: ``,
			ZH:  `收藏`,
			En:  `Bookmarked`,
			Ru:  `Добавить в закладки`,
			DE:  `mit Lesezeichen versehen`,
			FR:  `Favori`,
		}, {
			Num: 64,
			Img: ``,
			ZH:  `首页`,
			En:  `Home`,
			Ru:  `Дом`,
			DE:  `Startseite`,
			FR:  `Accueil`,
		}, {
			Num: 65,
			Img: ``,
			ZH:  `编辑器`,
			En:  `Editor`,
			Ru:  `Редактор`,
			DE:  `Editor`,
			FR:  `Éditeur`,
		}, {
			Num: 66,
			Img: ``,
			ZH:  `书签`,
			En:  `Bookmarks`,
			Ru:  `Закладки`,
			DE:  `Lesezeichen`,
			FR:  `Favoris`,
		}, {
			Num: 67,
			Img: ``,
			ZH:  `我的`,
			En:  `My Page`,
			Ru:  `Моя страница`,
			DE:  `Meine Seite`,
			FR:  `Ma page`,
		}, {
			Num: 68,
			Img: ``,
			ZH:  `关卡选择`,
			En:  `Select Stage`,
			Ru:  `Выберите этап`,
			DE:  `Phase wählen`,
			FR:  `Sélectionnez une étape`,
		}, {
			Num: 69,
			Img: ``,
			ZH:  `取消`,
			En:  `Cancel`,
			Ru:  `Отменить`,
			DE:  `abbrechen`,
			FR:  `Annuler`,
		}, {
			Num: 70,
			Img: ``,
			ZH:  `确定`,
			En:  `Confirm`,
			Ru:  `Подтверждить`,
			DE:  `bestätigen`,
			FR:  `Confirmer`,
		}, {
			Num: 71,
			Img: ``,
			ZH:  `我的书签`,
			En:  `My Bookmarks`,
			Ru:  `Мои закладки`,
			DE:  `Meine Lesezeichen`,
			FR:  `Mes favoris`,
		}, {
			Num: 72,
			Img: ``,
			ZH:  `攻略已下架`,
			En:  `Unlisted Guide`,
			Ru:  `Неопубликованное руководство`,
			DE:  `nicht aufgeführte Leitfaden`,
			FR:  `Guide non-publié`,
		}, {
			Num: 73,
			Img: ``,
			ZH:  `消息列表`,
			En:  `Notification Wall`,
			Ru:  `Центр уведомлений`,
			DE:  `Benachrichtigungstafel`,
			FR:  `Mur de notifications`,
		}, {
			Num: 74,
			Img: ``,
			ZH:  `消息`,
			En:  `Notification`,
			Ru:  `Уведомление`,
			DE:  `Benachrichtigung`,
			FR:  `Notification`,
		}, {
			Num: 75,
			Img: ``,
			ZH:  `WOR指南生成器`,
			En:  `WOR Guide Generator`,
			Ru:  `Создание руководств WOR`,
			DE:  `WOR Leitfaden-Generator`,
			FR:  `Générateur de guide WOR`,
		}, {
			Num: 76,
			Img: ``,
			ZH:  `标题`,
			En:  `Title`,
			Ru:  `Заголовок`,
			DE:  `Titel`,
			FR:  `Titre`,
		}, {
			Num: 77,
			Img: ``,
			ZH:  `每篇好的攻略都需要一个伟大的标题（20个字）`,
			En:  `It always takes a mighty title to make a good guide (20 characters max)`,
			Ru:  `Чтобы сделать хорошее руководство, всегда нужен громкий заголовок (максимум 20 символов)`,
			DE:  `Zu einem guten Leitfaden gehört immer ein aussagekräftiger Titel (maximal 20 Zeichen)`,
			FR:  `Il vous faut un titre percutant pour faire un bon guide (20 caractères max)`,
		}, {
			Num: 78,
			Img: ``,
			ZH:  `描述`,
			En:  `Description`,
			Ru:  `Описание`,
			DE:  `Beschreibung`,
			FR:  `Description`,
		}, {
			Num: 79,
			Img: ``,
			ZH: `简单描述你的好指南（200个字）
				ZH: `简单描述你的好指南（200个字）
例如: 1.主c英雄的战力区间
2.阵营小队激活效果
			En: `Describe your awesome guide in simple terms. (200 characters max)
				En: `Describe your awesome guide in simple terms. (200 characters max)
For example:
1. The recommended power (range) of main DPS hero.
2. The activition of faction effects
			Ru: `Опишите свое потрясающее руководство простыми словами. (Максимум 200 символов)
				Ru: `Опишите свое потрясающее руководство простыми словами. (Максимум 200 символов)
Например:
1. Рекомендуемая сила (диапазон) главного атакующего героя.
2. Активация эффектов фракции
			DE: `Beschreibe deinen fantastischen Leitfaden mit einfachen Worten. (maximal 200 Zeichen)
				DE: `Beschreibe deinen fantastischen Leitfaden mit einfachen Worten. (maximal 200 Zeichen)
Zum Beispiel:
1. Die empfohlene Stärke (Reichweite) des Haupt-DPS-Helden.
2. Die Aktivierung von Fraktionseffekten
			FR: `Décrivez votre super guide avec des termes simples. (200 caractères max)
				FR: `Décrivez votre super guide avec des termes simples. (200 caractères max)
Par exemple:
1. La puissance recommandée (environ) du héros DPS principal.
2. L'activation des effets de faction
		}, {
			Num: 80,
			Img: ``,
			ZH:  `排序`,
			En:  `Sort`,
			Ru:  `Сортировать`,
			DE:  `sortieren`,
			FR:  `Trier`,
		}, {
			Num: 81,
			Img: ``,
			ZH:  `全部重置`,
			En:  `Reset All`,
			Ru:  `Сбросить все`,
			DE:  `alles zurücksetzen`,
			FR:  `Tout réinitialiser`,
		}, {
			Num: 82,
			Img: ``,
			ZH:  `请配置你的英雄阵容，英雄分为可被上下位替换和不可替换英雄`,
			En:  `Please edit your hero deployments. Heroes are classified as replaceables (for better or worse) or non-replaceable heroes.`,
			Ru:  `Пожалуйста, отредактируйте расстановку героев. Герои классифицируются как заменяемые (к лучшему или к худшему) или незаменяемые герои.`,
			DE:  `Bitte bearbeite deine Heldeneinsätze. Helden werden in ersetzbare (für besser oder schlechter) und nicht ersetzbare Helden eingeteilt.`,
			FR:  `Veuillez modifier vos déploiements de héros. Les héros sont classifiés comme remplaçables (pour mieux ou moins bien) ou des héros non-remplaçables.`,
		}, {
			Num: 83,
			Img: ``,
			ZH:  `位置`,
			En:  `Location`,
			Ru:  `Расположение`,
			DE:  `Position`,
			FR:  `Emplacement`,
		}, {
			Num: 84,
			Img: ``,
			ZH:  `你可以选择已经预制好的地图，也可以自定义编辑关卡地图,然后按地图上的空格，选择你的英雄和他们的定位；`,
			En:  `You may select the premade maps or generate your own custome maps. Select the tiles on the map, and choose your heroes and their roles.`,
			Ru:  `Вы можете выбрать готовые карты или создать свои собственные пользовательские карты. Выберите плитки на карте и выберите героев и их роли.`,
			DE:  `Du kannst die vorgefertigten Karten auswählen oder deine eigenen Karten erstellen. Wähle die Kacheln auf der Karte und wähle deine Helden und ihre Rollen.`,
			FR:  `Vous pouvez sélectionner les cartes prédéfinies ou générer vos propres cartes personnalisées. Sélectionnez les cases sur la carte et choisissez vos héros et leurs rôles.`,
		}, {
			Num: 85,
			Img: ``,
			ZH:  `选择您的地图`,
			En:  `Select Your Map`,
			Ru:  `Выберите карту`,
			DE:  `Wähle deine Karte`,
			FR:  `Sélectionnez votre carte`,
		}, {
			Num: 86,
			Img: ``,
			ZH:  `自定义地图`,
			En:  `Custom Map`,
			Ru:  `Пользовательская карта`,
			DE:  `Benutzerdefinierte Karte`,
			FR:  `Carte personnalisée`,
		}, {
			Num: 87,
			Img: ``,
			ZH: `1.在选项卡之间切换，为不同的战斗阶段设置英雄
				ZH: `1.在选项卡之间切换，为不同的战斗阶段设置英雄
			En: `1. Switch between option cards, deploying different heroes during different battle stages.
				En: `1. Switch between option cards, deploying different heroes during different battle stages.
			Ru: `1. Переключайтесь между вариантами карт, размещая разных героев на разных этапах битвы. 
				Ru: `1. Переключайтесь между вариантами карт, размещая разных героев на разных этапах битвы. 
			DE: `1. Wechsle zwischen den Optionskarten und setze verschiedene Helden in verschiedenen Kampfphasen ein.
				DE: `1. Wechsle zwischen den Optionskarten und setze verschiedene Helden in verschiedenen Kampfphasen ein.
			FR: `1. Basculez entre les cartes d'option en déployant différents héros au cours de différentes étapes.
				FR: `1. Basculez entre les cartes d'option en déployant différents héros au cours de différentes étapes.
		}, {
			Num: 88,
			Img: ``,
			ZH:  `line up1 的描述内容`,
			En:  `Description content of Line Up 1`,
			Ru:  `Описание содержания 1 состава`,
			DE:  `Beschreibung des Inhalts der 1. Aufstellung`,
			FR:  `Contenu de la description de l'équipe 1`,
		}, {
			Num: 89,
			Img: ``,
			ZH:  `配置英雄装备`,
			En:  `Hero Gear Loadout`,
			Ru:  `Комплект снаряжения героя`,
			DE:  `Heldenausrüstung`,
			FR:  `Configuration d'équipement de héros`,
		}, {
			Num: 90,
			Img: ``,
			ZH:  `简述你的配装思路（100个字符）`,
			En:  `Describe your loadout ideas in simple terms. (100 characters max)`,
			Ru:  `Опишите ваши идеи по снаряжению простыми словами. (Максимум 100 символов)`,
			DE:  `Beschreibe deine Ausrüstungsideen in einfachen Worten. (maximal 100 Zeichen)`,
			FR:  `Décrivez vos idées de configuration d'équipements avec des termes simples. (100 caractères max)`,
		}, {
			Num: 91,
			Img: ``,
			ZH:  `英雄`,
			En:  `Hero`,
			Ru:  `Герой`,
			DE:  `Held`,
			FR:  `Héros`,
		}, {
			Num: 92,
			Img: ``,
			ZH:  `装备效果`,
			En:  `Gear Effect`,
			Ru:  `Эффект снаряжения`,
			DE:  `Ausrüstungseffekt`,
			FR:  `Effet d'équipement`,
		}, {
			Num: 93,
			Img: ``,
			ZH:  `神器`,
			En:  `Artifact`,
			Ru:  `Артефакт`,
			DE:  `Artefakt`,
			FR:  `Artéfact`,
		}, {
			Num: 94,
			Img: ``,
			ZH:  `点击英雄头像修改配置`,
			En:  `Click the hero portrait to change the loadout`,
			Ru:  `Нажмите на портрет героя, чтобы сменить снаряжение.`,
			DE:  `Klicke auf das Heldenporträt, um die Ausrüstung zu ändern`,
			FR:  `Cliquez sur le portrait du héros pour changer l'équipement`,
		}, {
			Num: 95,
			Img: ``,
			ZH:  `专属神器`,
			En:  `Exclusive Artifact`,
			Ru:  `Эксклюзивный Артефакт`,
			DE:  `Exklusives Artefakt`,
			FR:  `Artéfact exclusif`,
		}, {
			Num: 96,
			Img: ``,
			ZH:  `通用神器`,
			En:  `Generic Artifact`,
			Ru:  `Общий Артефакт`,
			DE:  `Allgemeines Artefakt`,
			FR:  `Artéfact générique`,
		}, {
			Num: 97,
			Img: ``,
			ZH:  `(选填)请简单描述装备套装，如主副属性、强化目标等(50个字符)`,
			En:  `(Optional) Please describe the gear loadout in simple terms, such as main/sub stats and the optimal enchaned value. (50 characters max)`,
			Ru:  `(Необязательно) Опишите комплект снаряжения простыми словами, например, основные/дополнительные характеристики и оптимальное улучшенное значение. (Максимум 50 символов)`,
			DE:  `(Optional) Bitte beschreibe die Ausrüstung in einfachen Worten, z. B. Haupt- und Nebenwerte und den optimalen Verzauberungswert. (maximal 50 Zeichen)`,
			FR:  `(Optionnel) Veuillez décrire la configuration des équipements en termes simples, tels que les attributs principaux/secondaires et la valeur de renforcement optimale. (50 caractères maximum)`,
		}, {
			Num: 98,
			Img: ``,
			ZH:  `添加英雄`,
			En:  `Add A New Hero`,
			Ru:  `Добавить нового героя`,
			DE:  `Einen neuen Helden hinzufügen`,
			FR:  `Ajouter un nouveau héros`,
		}, {
			Num: 99,
			Img: ``,
			ZH:  `攻略视频`,
			En:  `Video Guide`,
			Ru:  `Видеоруководство`,
			DE:  `Video-Leitfaden`,
			FR:  `Guide vidéo`,
		}, {
			Num: 100,
			Img: ``,
			ZH:  `(选填)攻略视频链接`,
			En:  `(Optional) Link to the video guide`,
			Ru:  `(Необязательно) Ссылка на видеоруководство`,
			DE:  `(Optional) Link zum Video-Leitfaden`,
			FR:  `(Optionnel) Lien de la vidéo guide`,
		}, {
			Num: 101,
			Img: ``,
			ZH:  `预览`,
			En:  `Preview`,
			Ru:  `Предварительный просмотр`,
			DE:  `Vorschau`,
			FR:  `Aperçu`,
		}, {
			Num: 102,
			Img: ``,
			ZH:  `发布`,
			En:  `Publish`,
			Ru:  `Публиковать`,
			DE:  `veröffentlichen`,
			FR:  `Publier`,
		}, {
			Num: 103,
			Img: ``,
			ZH:  `保存为长图`,
			En:  `Save As Panorama`,
			Ru:  `Сохранить как панораму`,
			DE:  `Als Panorama speichern`,
			FR:  `Sauvegarder en tant que panorama`,
		}, {
			Num: 104,
			Img: ``,
			ZH:  `存为草稿`,
			En:  `Save As Draft`,
			Ru:  `Сохранить как черновик`,
			DE:  `Als Entwurf speichern`,
			FR:  `Sauvegarder en tant que brouillon`,
		}, {
			Num: 105,
			Img: ``,
			ZH:  `上阵英雄选择`,
			En:  `Select Deployed Hero`,
			Ru:  `Выберите размещенного героя`,
			DE:  `Eingesetzten Helden auswählen`,
			FR:  `Sélectionnez le héros déployé`,
		}, {
			Num: 106,
			Img: ``,
			ZH:  `选择上阵英雄（必选）`,
			En:  `Select The Hero To Deploy (Required)`,
			Ru:  `Выберите героя для размещения (обязательно)`,
			DE:  `Wähle den einzusetzenden Helden (erforderlich)`,
			FR:  `Sélectionnez le héros à déployer (Obligatoire)`,
		}, {
			Num: 107,
			Img: ``,
			ZH:  `选择英雄方向（必选）`,
			En:  `Select The Direction Hero Will Be Facing (Required)`,
			Ru:  `Выберите направление, куда будет смотреть герой (обязательно)`,
			DE:  `Wähle die Richtung, in die der Held blickt (erforderlich)`,
			FR:  `Sélectionnez la direction vers laquelle le héros fera face (Obligatoire)`,
		}, {
			Num: 108,
			Img: ``,
			ZH:  `向上`,
			En:  `Upwards`,
			Ru:  `Вверх`,
			DE:  `Nach oben`,
			FR:  `Vers le haut`,
		}, {
			Num: 109,
			Img: ``,
			ZH:  `向下`,
			En:  `Downwards`,
			Ru:  `Вниз`,
			DE:  `Nach unten`,
			FR:  `Vers le bas`,
		}, {
			Num: 110,
			Img: ``,
			ZH:  `向左`,
			En:  `Leftwards`,
			Ru:  `Влево`,
			DE:  `Nach links`,
			FR:  `Vers la gauche`,
		}, {
			Num: 111,
			Img: ``,
			ZH:  `向右`,
			En:  `Rightwards`,
			Ru:  `Вправо`,
			DE:  `Nach rechts`,
			FR:  `Vers la droite`,
		}, {
			Num: 112,
			Img: ``,
			ZH:  `上阵顺序（必选）`,
			En:  `Sequence Of Deploying (Required)`,
			Ru:  `Последовательность размещения (обязательно)`,
			DE:  `Reihenfolge des Einsatzes (erforderlich)`,
			FR:  `Ordre de déploiement (Obligatoire)`,
		}, {
			Num: 113,
			Img: ``,
			ZH:  `请选择上阵英雄序号`,
			En:  `Please assign the sequence number for each deployed hero`,
			Ru:  `Пожалуйста, присвойте порядковый номер каждому размещенному герою.`,
			DE:  `Bitte vergebe die Nummer der Reihenfolge für jeden eingesetzten Helden`,
			FR:  `Veuillez préciser dans quel ordre les héros seront déployés en leur donnant un nombre`,
		}, {
			Num: 114,
			Img: ``,
			ZH:  `确定`,
			En:  `Confirm`,
			Ru:  `Подтвердить`,
			DE:  `bestätigen`,
			FR:  `Confirmer`,
		}, {
			Num: 115,
			Img: ``,
			ZH:  `上阵英雄配置`,
			En:  `Loadout of Deployed Heroes`,
			Ru:  `Выберите размещенных героев`,
			DE:  `Ausrüstung der eingesetzten Helden`,
			FR:  `Configuration d'équipements des héros déployés`,
		}, {
			Num: 116,
			Img: ``,
			ZH:  `选择配置英雄（必选）`,
			En:  `Select The Heroes To Edit loadout`,
			Ru:  `Выберите героев для редактирования снаряжения`,
			DE:  `Wähle die Helden um die Ausrüstung zu bearbeiten`,
			FR:  `Sélectionner les héros pour modifier les équipements`,
		}, {
			Num: 117,
			Img: ``,
			ZH:  `已配置`,
			En:  `Geared-Up`,
			Ru:  `Экипировать`,
			DE:  `aufgerüstet`,
			FR:  `Équipé`,
		}, {
			Num: 118,
			Img: ``,
			ZH:  `选择英雄装备效果（非必选）`,
			En:  `Select The Effect Of Hero's Gears`,
			Ru:  `Выберите эффект снаряжения героя`,
			DE:  `Wählen den Effekt für die Ausrüstung des Helden`,
			FR:  `Sélectionner les effets des équipements de héros`,
		}, {
			Num: 119,
			Img: ``,
			ZH:  `武器胸甲`,
			En:  `Weapon/Breastplate (2-piece set)`,
			Ru:  `Оружие/Нагрудник (комплект из 2 предметов)`,
			DE:  `Waffe/Brustpanzer (2-teiliges Set)`,
			FR:  `Arme/Plastron (Ensemble 2 pièces)`,
		}, {
			Num: 120,
			Img: ``,
			ZH:  `首饰`,
			En:  `Bangle/Amulet/Ring (3-piece set)`,
			Ru:  `Браслет/Амулет/Кольцо (набор из 3 предметов)`,
			DE:  `Armreif/Amulett/Ring (3-teiliges Set)`,
			FR:  `Bracelet/Amulette/Anneau (Ensemble 3 pièces)`,
		}, {
			Num: 121,
			Img: ``,
			ZH:  `选择英雄神器（非必选）`,
			En:  `Select Hero Artifacts`,
			Ru:  `Выберите артефакты героя`,
			DE:  `Heldenartefakte auswählen`,
			FR:  `Sélectionner les artéfacts des héros`,
		}, {
			Num: 122,
			Img: ``,
			ZH:  `先选择的为首选神器后选择的为次选神器`,
			En:  `Click once on a artifact to define it as the optimal artifact. Then click on another artifact to define it as the second best artfiact choice.`,
			Ru:  `Нажмите один раз на артефакт, чтобы определить его как оптимальный артефакт. Затем нажмите на другой артефакт, чтобы определить его как второй лучший выбор артефакта.`,
			DE:  `Klicke einmal auf ein Artefakt, um es als das optimale Artefakt zu definieren. Klicke dann auf ein anderes Artefakt, um es als zweitbeste Artefaktwahl zu definieren.`,
			FR:  `Cliquez une fois sur un artefact pour le définir comme artefact optimal. Cliquez ensuite sur un autre artefact pour le définir comme le deuxième meilleur choix d'artefact.`,
		}, {
			Num: 123,
			Img: ``,
			ZH:  `法师`,
			En:  `Mage`,
			Ru:  `Маг`,
			DE:  `Magier`,
			FR:  `Mage`,
		}, {
			Num: 124,
			Img: ``,
			ZH:  `射手`,
			En:  `Marksman`,
			Ru:  `Стрелок`,
			DE:  `Schütze`,
			FR:  `Tireur`,
		}, {
			Num: 125,
			Img: ``,
			ZH:  `守护者`,
			En:  `Defender`,
			Ru:  `Защитник`,
			DE:  `Verteidiger`,
			FR:  `Gardien`,
		}, {
			Num: 126,
			Img: ``,
			ZH:  `医师`,
			En:  `Healer`,
			Ru:  `Целитель`,
			DE:  `Heiler`,
			FR:  `Guérisseur`,
		}, {
			Num: 127,
			Img: ``,
			ZH:  `战士`,
			En:  `Fighter`,
			Ru:  `Боец`,
			DE:  `Kämpfer`,
			FR:  `Combattant`,
		}, {
			Num: 128,
			Img: ``,
			ZH:  `3星`,
			En:  `3 Star`,
			Ru:  `3 звезды`,
			DE:  `3-Sterne`,
			FR:  `3 étoiles`,
		}, {
			Num: 129,
			Img: ``,
			ZH:  `4星`,
			En:  `4 Star`,
			Ru:  `4 звезды`,
			DE:  `4-Sterne`,
			FR:  `4 étoiles`,
		}, {
			Num: 130,
			Img: ``,
			ZH:  `5星`,
			En:  `5 Star`,
			Ru:  `5 звезд`,
			DE:  `5-Sterne`,
			FR:  `5 étoiles`,
		}, {
			Num: 131,
			Img: ``,
			ZH:  `保存`,
			En:  `Save`,
			Ru:  `Сохранить`,
			DE:  `speichern`,
			FR:  `Sauvegarder`,
		}, {
			Num: 132,
			Img: ``,
			ZH:  `提示`,
			En:  `Notification`,
			Ru:  `Уведомление`,
			DE:  `Benachrichtigung`,
			FR:  `Notification`,
		}, {
			Num: 133,
			Img: ``,
			ZH:  `当前内容未保存`,
			En:  `Current edits have not been saved`,
			Ru:  `Текущие изменения не были сохранены`,
			DE:  `Die aktuellen Bearbeitungen wurden nicht gespeichert`,
			FR:  `Les dernières modifications n'ont pas été sauvegardées`,
		}, {
			Num: 134,
			Img: ``,
			ZH:  `返回`,
			En:  `Return`,
			Ru:  `Вернуть`,
			DE:  `Rückkehren`,
			FR:  `Retour`,
		}, {
			Num: 135,
			Img: ``,
			ZH:  `保存草稿箱`,
			En:  `Save To Draft`,
			Ru:  `Сохранить в черновике`,
			DE:  `Im Entwurf speichern`,
			FR:  `Sauvegarder le brouillon`,
		}, {
			Num: 136,
			Img: ``,
			ZH:  `选择英雄`,
			En:  `Select Hero`,
			Ru:  `Выбрать героя`,
			DE:  `Helden wählen`,
			FR:  `Sélectionner le héros`,
		}, {
			Num: 137,
			Img: ``,
			ZH:  `保存`,
			En:  `Save`,
			Ru:  `Сохранить`,
			DE:  `speichern`,
			FR:  `Sauvegarder`,
		}, {
			Num: 138,
			Img: ``,
			ZH:  `英雄池`,
			En:  `Hero Pool`,
			Ru:  `Пул героев`,
			DE:  `Heldenpool`,
			FR:  `Pool de héros`,
		}, {
			Num: 139,
			Img: ``,
			ZH:  `点击英雄头像来上阵该英雄`,
			En:  `Click on the hero portrait to deploy this hero`,
			Ru:  `Нажмите на портрет героя, чтобы разместить этого героя`,
			DE:  `Klicke auf das Heldenporträt, um diesen Helden einzusetzen`,
			FR:  `Cliquez sur le portrait du héros pour le déployer`,
		}, {
			Num: 140,
			Img: ``,
			ZH:  `不可替代英雄`,
			En:  `Non-replaceable heroes`,
			Ru:  `Незаменимые герои`,
			DE:  `Unersetzbare Helden`,
			FR:  `Héros non-remplaçables`,
		}, {
			Num: 141,
			Img: ``,
			ZH:  `法师`,
			En:  `Mage`,
			Ru:  `Маг`,
			DE:  `Magier`,
			FR:  `Mage`,
		}, {
			Num: 142,
			Img: ``,
			ZH:  `射手`,
			En:  `Marksman`,
			Ru:  `Стрелок`,
			DE:  `Schütze`,
			FR:  `Tireur`,
		}, {
			Num: 143,
			Img: ``,
			ZH:  `守护者`,
			En:  `Defender`,
			Ru:  `Защитник`,
			DE:  `Verteidiger`,
			FR:  `Gardien`,
		}, {
			Num: 144,
			Img: ``,
			ZH:  `医师`,
			En:  `Healer`,
			Ru:  `Целитель`,
			DE:  `Heiler`,
			FR:  `Guérisseur`,
		}, {
			Num: 145,
			Img: ``,
			ZH:  `战士`,
			En:  `Fighter`,
			Ru:  `Боец`,
			DE:  `Kämpfer`,
			FR:  `Combattant`,
		}, {
			Num: 146,
			Img: ``,
			ZH:  `已选择`,
			En:  `Selected`,
			Ru:  `Выбрано`,
			DE:  `Ausgewählt`,
			FR:  `Sélectionné`,
		}, {
			Num: 147,
			Img: ``,
			ZH:  `可被替代英雄`,
			En:  `Replaceable heroes`,
			Ru:  `Заменяемые герои`,
			DE:  `Ersetzbare Helden`,
			FR:  `Héros remplaçables`,
		}, {
			Num: 148,
			Img: ``,
			ZH:  `提示`,
			En:  `Notification`,
			Ru:  `Уведомление`,
			DE:  `Benachrichtigung`,
			FR:  `Notification`,
		}, {
			Num: 149,
			Img: ``,
			ZH:  `是否保存当前阵容`,
			En:  `Save the current formation?`,
			Ru:  `Сохранить текущую формацию?`,
			DE:  `Die aktuelle Formation speichern?`,
			FR:  `Sauvegarder la formation actuelle ?`,
		}, {
			Num: 150,
			Img: ``,
			ZH:  `返回`,
			En:  `Return`,
			Ru:  `Вернуть`,
			DE:  `Rückkehren`,
			FR:  `Retour`,
		}, {
			Num: 151,
			Img: ``,
			ZH:  `保存`,
			En:  `Save`,
			Ru:  `Сохранить`,
			DE:  `Speichern`,
			FR:  `Sauvegarder`,
		}, {
			Num: 152,
			Img: ``,
			ZH:  `关卡地图选择`,
			En:  `Map Stage Selection`,
			Ru:  `Выбор этапа карты`,
			DE:  `Auswahl der Kartenstufe`,
			FR:  `Sélection de la carte`,
		}, {
			Num: 153,
			Img: ``,
			ZH:  `选择关卡`,
			En:  `Select Stage`,
			Ru:  `Выберите этап`,
			DE:  `Phase wählen`,
			FR:  `Sélectionner l'étape`,
		}, {
			Num: 154,
			Img: ``,
			ZH:  `关卡地图选择`,
			En:  `Map Stage Selection`,
			Ru:  `Выбор этапа карты`,
			DE:  `Auswahl der Kartenstufe`,
			FR:  `Sélection de la carte`,
		}, {
			Num: 155,
			Img: ``,
			ZH:  `选择关卡`,
			En:  `Select Stage`,
			Ru:  `Выберите этап`,
			DE:  `Phase wählen`,
			FR:  `Sélectionner l'étape`,
		}, {
			Num: 156,
			Img: ``,
			ZH:  `请选择上下位英雄`,
			En:  `Please Select The Replaceable Heroes`,
			Ru:  `Пожалуйста, выберите заменяемых героев`,
			DE:  `Bitte wähle die ersetzbaren Helden`,
			FR:  `Veuillez sélectionner les héros remplaçables`,
		}, {
			Num: 157,
			Img: ``,
			ZH:  `英雄池`,
			En:  `Hero Pool`,
			Ru:  `Пул героев`,
			DE:  `Heldenpool`,
			FR:  `Pool de héros`,
		}, {
			Num: 158,
			Img: ``,
			ZH:  `点击英雄头像来上阵该英雄`,
			En:  `Click on the hero portrait to deploy this hero`,
			Ru:  `Нажмите на портрет героя, чтобы разместить этого героя`,
			DE:  `Klicke auf das Heldenporträt, um diesen Helden einzusetzen`,
			FR:  `Cliquez sur le portrait du héros pour le déployer`,
		}, {
			Num: 159,
			Img: ``,
			ZH:  `可被替代英雄`,
			En:  `Replaceable heroes`,
			Ru:  `Заменяемые герои`,
			DE:  `Ersetzbare Helden`,
			FR:  `Héros remplaçables`,
		}, {
			Num: 160,
			Img: ``,
			ZH:  `已选择`,
			En:  `Selected`,
			Ru:  `Выбрано`,
			DE:  `ausgewählt`,
			FR:  `Sélectionné`,
		}, {
			Num: 161,
			Img: ``,
			ZH:  `保存`,
			En:  `Save`,
			Ru:  `Сохранить`,
			DE:  `speichern`,
			FR:  `Sauvegarder`,
		}, {
			Num: 162,
			Img: ``,
			ZH:  `WOR地图编辑器`,
			En:  `WOR Map Generator`,
			Ru:  `Генератор карт WOR`,
			DE:  `WOR Karten-Generator`,
			FR:  `Générateur de carte WOR`,
		}, {
			Num: 163,
			Img: ``,
			ZH:  `选择绘制的地图关卡`,
			En:  `Select the map stage to edit`,
			Ru:  `Выберите этап карты для редактирования`,
			DE:  `Wähle die zu bearbeitende Kartenstufe`,
			FR:  `Sélectionnez la carte à modifier`,
		}, {
			Num: 164,
			Img: ``,
			ZH:  `请选择地图的长度`,
			En:  `Please define the length of map`,
			Ru:  `Пожалуйста, определите длину карты`,
			DE:  `Bitte definiere die Länge der Karte`,
			FR:  `Veuillez définir la longueur de la carte`,
		}, {
			Num: 165,
			Img: ``,
			ZH:  `请选择地图的宽度`,
			En:  `Please define the width of map`,
			Ru:  `Пожалуйста, определите ширину карты`,
			DE:  `Bitte definiere die Breite der Karte`,
			FR:  `Veuillez définir la largeur de la carte`,
		}, {
			Num: 166,
			Img: ``,
			ZH:  `保存`,
			En:  `Save`,
			Ru:  `Сохранить`,
			DE:  `Speichern`,
			FR:  `Sauvegarder`,
		}, {
			Num: 167,
			Img: ``,
			ZH:  `确定`,
			En:  `Confirm`,
			Ru:  `Подтвердить`,
			DE:  `Bestätigen`,
			FR:  `Confirmer`,
		}, {
			Num: 168,
			Img: ``,
			ZH:  `取消`,
			En:  `Cancel`,
			Ru:  `Отменить`,
			DE:  `Abbrechen`,
			FR:  `Annuler`,
		}, {
			Num: 169,
			Img: ``,
			ZH:  `视频链接`,
			En:  `Link To The Video`,
			Ru:  `Ссылка на видео`,
			DE:  `Link zum Video`,
			FR:  `Lien vers la vidéo`,
		}, {
			Num: 170,
			Img: ``,
			ZH:  `地图位置`,
			En:  `Map Location`,
			Ru:  `Расположение на карте`,
			DE:  `Karten Position`,
			FR:  `Emplacement de la carte`,
		}, {
			Num: 171,
			Img: ``,
			ZH:  `英雄装备攻略`,
			En:  `Hero Gear Guide`,
			Ru:  `Руководство по снаряжению героя`,
			DE:  `Heldenausrüstungs Leitfaden`,
			FR:  `Guide d'équipement de héros`,
		}, {
			Num: 172,
			Img: ``,
			ZH:  `制作我自己的攻略`,
			En:  `Create My Own Guide`,
			Ru:  `Создать собственное руководство`,
			DE:  `Meinen eigenen Leitfaden erstellen`,
			FR:  `Créer mon propre guide`,
		}, {
			Num: 173,
			Img: ``,
			ZH:  `举报反馈`,
			En:  `Report Feedback`,
			Ru:  `Сообщить об ошибке`,
			DE:  `Feedback melden`,
			FR:  `Donnez votre avis`,
		}, {
			Num: 174,
			Img: ``,
			ZH:  `作者名`,
			En:  `Creator`,
			Ru:  `Создатель`,
			DE:  `Ersteller`,
			FR:  `Créateur`,
		}, {
			Num: 175,
			Img: ``,
			ZH:  `攻略标题`,
			En:  `Guide Title`,
			Ru:  `Название руководства`,
			DE:  `Leitfaden des Titels`,
			FR:  `Titre du guide`,
		}, {
			Num: 176,
			Img: ``,
			ZH:  `举报原因（必填）`,
			En:  `Report Reason`,
			Ru:  `Причина отчета`,
			DE:  `Grund für die Meldung`,
			FR:  `Raison du signalement`,
		}, {
			Num: 177,
			Img: ``,
			ZH:  `请输入原因`,
			En:  `Please input the reason(s)`,
			Ru:  `Пожалуйста, введите причину(ы)`,
			DE:  `Bitte gib die Gründe ein`,
			FR:  `Veuillez saisir la ou les raisons`,
		}, {
			Num: 178,
			Img: ``,
			ZH:  `图片证据（最多上传9张）`,
			En:  `Image Evidence(s) (9 pieces max)`,
			Ru:  `Подтверждающие изображения (макс. 9 шт.)`,
			DE:  `Bildnachweis(e) (maximal 9 Stück)`,
			FR:  `Image(s) de preuve (9 images max)`,
		}, {
			Num: 179,
			Img: ``,
			ZH:  `提交`,
			En:  `Submit`,
			Ru:  `Представить на рассмотрение`,
			DE:  `Einreichen`,
			FR:  `Envoyer`,
		}, {
			Num: 180,
			Img: ``,
			ZH:  `提示`,
			En:  `Notification`,
			Ru:  `Уведомление`,
			DE:  `Benachrichtigung`,
			FR:  `Notification`,
		}, {
			Num: 181,
			Img: ``,
			ZH:  `即将离开本页面，请注意账号安全`,
			En:  `You are about to leave this page, please keep your account save`,
			Ru:  `Вы собираетесь покинуть эту страницу, пожалуйста, сохраните свою учетную запись`,
			DE:  `Du bist dabei diese Seite zu verlassen, bitte bewahre dein Konto sicher auf`,
			FR:  `Vous êtes sur le point de quitter cette page, veuillez garder votre compte à jour`,
		}, {
			Num: 182,
			Img: ``,
			ZH:  `继续访问`,
			En:  `Continue to visit`,
			Ru:  `Продолжить посещение`,
			DE:  `Weiter besuchen`,
			FR:  `Continuez à visiter`,
		}, {
			Num: 183,
			Img: ``,
			ZH:  `提示`,
			En:  `Notification`,
			Ru:  `Уведомление`,
			DE:  `Benachrichtigung`,
			FR:  `Notification`,
		}, {
			Num: 184,
			Img: ``,
			ZH:  `提交成功，我们将尽快处理`,
			En:  `Submit`,
			Ru:  `Представить на рассмотрение`,
			DE:  `einreichen`,
			FR:  `Envoyer`,
		}, {
			Num: 185,
			Img: ``,
			ZH:  `我的攻略`,
			En:  `My Guides`,
			Ru:  `Мои руководства`,
			DE:  `Meine Leitfäden`,
			FR:  `Mes guides`,
		}, {
			Num: 186,
			Img: ``,
			ZH:  `已下架`,
			En:  `Unlisted`,
			Ru:  `Не действительно`,
			DE:  `nicht aufgeführt`,
			FR:  `Non-publié`,
		}, {
			Num: 187,
			Img: ``,
			ZH:  `草稿箱`,
			En:  `Drafts`,
			Ru:  `Черновики`,
			DE:  `Entwürfe`,
			FR:  `Brouillons`,
		}, {
			Num: 188,
			Img: ``,
			ZH:  `下架`,
			En:  `Unlist`,
			Ru:  `Удалить из списка`,
			DE:  `abwählen`,
			FR:  `Dépublier`,
		}, {
			Num: 189,
			Img: ``,
			ZH:  `删除`,
			En:  `Delete`,
			Ru:  `Удалить`,
			DE:  `löschen`,
			FR:  `Supprimer`,
		}, {
			Num: 190,
			Img: ``,
			ZH:  `提示`,
			En:  `Notification`,
			Ru:  `Уведомление`,
			DE:  `Benachrichtigung`,
			FR:  `Notification`,
		}, {
			Num: 191,
			Img: ``,
			ZH:  `确定删除该内容吗？`,
			En:  `Are you sure you want to remove this content?`,
			Ru:  `Вы уверены, что хотите удалить этот контент?`,
			DE:  `Bist du sicher, dass du diesen Inhalt entfernen möchtest?`,
			FR:  `Êtes-vous sûr de vouloir supprimer ce contenu ?`,
		}, {
			Num: 192,
			Img: ``,
			ZH:  `确定`,
			En:  `Confirm`,
			Ru:  `Подтверждить`,
			DE:  `bestätigen`,
			FR:  `Confirmer`,
		}, {
			Num: 193,
			Img: ``,
			ZH:  `取消`,
			En:  `Cancel`,
			Ru:  `Отменить`,
			DE:  `abbrechen`,
			FR:  `Annuler`,
		}, {
			Num: 194,
			Img: ``,
			ZH:  `建议反馈`,
			En:  `Suggestion Feedback`,
			Ru:  `Предоставить Обратная связь`,
			DE:  `Feedback zu Vorschlägen`,
			FR:  `Faire une suggestion`,
		}, {
			Num: 195,
			Img: ``,
			ZH:  `编辑`,
			En:  `Edit`,
			Ru:  `Редактировать`,
			DE:  `bearbeiten`,
			FR:  `Modifier`,
		}, {
			Num: 196,
			Img: ``,
			ZH:  `建议（必填）`,
			En:  `Suggestions (Required)`,
			Ru:  `Предложения (обязательно)`,
			DE:  `Vorschläge (erforderlich)`,
			FR:  `Suggestions (Requis)`,
		}, {
			Num: 197,
			Img: ``,
			ZH:  `请输入内容`,
			En:  `Please input content`,
			Ru:  `Пожалуйста, введите содержимое`,
			DE:  `Bitte Inhalt eingeben`,
			FR:  `Veuillez saisir du contenu`,
		}, {
			Num: 198,
			Img: ``,
			ZH:  `上传图片（最多上传9张）`,
			En:  `Upload Image(s) (9 pieces max)`,
			Ru:  `Загрузить изображение(я) (максимум 9 шт.)`,
			DE:  `Bild(er) hochladen (maximal 9 Stück)`,
			FR:  `Télécharger une ou des images (9 images max)`,
		}, {
			Num: 199,
			Img: ``,
			ZH:  `装备副本I`,
			En:  `Gear Raid I`,
			Ru:  `Рейд для снаряжения I`,
			DE:  `Ausrüstungsraid I`,
			FR:  `Raid d'équipement I`,
		}, {
			Num: 200,
			Img: ``,
			ZH:  `装备副本II`,
			En:  `Gear Raid II`,
			Ru:  `Рейд для снаряжения II`,
			DE:  `Ausrüstungsraid II`,
			FR:  `Raid d'équipement II`,
		}, {
			Num: 201,
			Img: ``,
			ZH:  `装备副本III`,
			En:  `Gear Raid III`,
			Ru:  `Рейд для снаряжения III`,
			DE:  `Ausrüstungsraid III`,
			FR:  `Raid d'équipement III`,
		}, {
			Num: 202,
			Img: ``,
			ZH:  `神器材料副本`,
			En:  `Artifact Material Raid`,
			Ru:  `Рейд для материалов артефактов`,
			DE:  `Artefaktmaterial-Raid`,
			FR:  `Raid de matériaux d'artefacts`,
		}, {
			Num: 203,
			Img: ``,
			ZH:  `装备地下城I`,
			En:  `Gear Dungeon I`,
			Ru:  `Подземелье снаряжения I`,
			DE:  `Ausrüstungs-Dungeon I`,
			FR:  `Donjon d'équipement I`,
		}, {
			Num: 204,
			Img: ``,
			ZH:  `装备地下城II`,
			En:  `Gear Dungeon II`,
			Ru:  `Подземелье снаряжения II`,
			DE:  `Ausrüstungs-Dungeon II`,
			FR:  `Donjon d'équipement II`,
		}, {
			Num: 205,
			Img: ``,
			ZH:  `关卡1`,
			En:  `Stage 1`,
			Ru:  `Этап 1`,
			DE:  `Stufe 1`,
			FR:  `Étape 1`,
		}, {
			Num: 206,
			Img: ``,
			ZH:  `关卡2`,
			En:  `Stage 2`,
			Ru:  `Этап 2`,
			DE:  `Stufe 2`,
			FR:  `Étape 2`,
		}, {
			Num: 207,
			Img: ``,
			ZH:  `关卡3`,
			En:  `Stage 3`,
			Ru:  `Этап 3`,
			DE:  `Stufe 3`,
			FR:  `Étape 3`,
		}, {
			Num: 208,
			Img: ``,
			ZH:  `关卡4`,
			En:  `Stage 4`,
			Ru:  `Этап 4`,
			DE:  `Stufe 4`,
			FR:  `Étape 4`,
		}, {
			Num: 209,
			Img: ``,
			ZH:  `关卡5`,
			En:  `Stage 5`,
			Ru:  `Этап 5`,
			DE:  `Stufe 5`,
			FR:  `Étape 5`,
		}, {
			Num: 210,
			Img: ``,
			ZH:  `关卡6`,
			En:  `Stage 6`,
			Ru:  `Этап 6`,
			DE:  `Stufe 6`,
			FR:  `Étape 6`,
		}, {
			Num: 211,
			Img: ``,
			ZH:  `关卡7`,
			En:  `Stage 7`,
			Ru:  `Этап 7`,
			DE:  `Stufe 7`,
			FR:  `Étape 7`,
		}, {
			Num: 212,
			Img: ``,
			ZH:  `关卡8`,
			En:  `Stage 8`,
			Ru:  `Этап 8`,
			DE:  `Stufe 8`,
			FR:  `Étape 8`,
		}, {
			Num: 213,
			Img: ``,
			ZH:  `关卡9`,
			En:  `Stage 9`,
			Ru:  `Этап 9`,
			DE:  `Stufe 9`,
			FR:  `Étape 9`,
		}, {
			Num: 214,
			Img: ``,
			ZH:  `关卡10`,
			En:  `Stage 10`,
			Ru:  `Этап 10`,
			DE:  `Stufe 10`,
			FR:  `Étape 10`,
		}, {
			Num: 215,
			Img: ``,
			ZH:  `关卡11`,
			En:  `Stage 11`,
			Ru:  `Этап 11`,
			DE:  `Stufe 11`,
			FR:  `Étape 11`,
		}, {
			Num: 216,
			Img: ``,
			ZH:  `关卡12`,
			En:  `Stage 12`,
			Ru:  `Этап 12`,
			DE:  `Stufe 12`,
			FR:  `Étape 12`,
		}, {
			Num: 217,
			Img: ``,
			ZH:  `关卡13`,
			En:  `Stage 13`,
			Ru:  `Этап 13`,
			DE:  `Stufe 13`,
			FR:  `Étape 13`,
		}, {
			Num: 218,
			Img: ``,
			ZH:  `关卡14`,
			En:  `Stage 14`,
			Ru:  `Этап 14`,
			DE:  `Stufe 14`,
			FR:  `Étape 14`,
		}, {
			Num: 219,
			Img: ``,
			ZH:  `关卡15`,
			En:  `Stage 15`,
			Ru:  `Этап 15`,
			DE:  `Stufe 15`,
			FR:  `Étape 15`,
		}, {
			Num: 220,
			Img: ``,
			ZH:  `关卡16`,
			En:  `Stage 16`,
			Ru:  `Этап 16`,
			DE:  `Stufe 16`,
			FR:  `Étape 16`,
		}, {
			Num: 221,
			Img: ``,
			ZH:  `关卡17`,
			En:  `Stage 17`,
			Ru:  `Этап 17`,
			DE:  `Stufe 17`,
			FR:  `Étape 17`,
		}, {
			Num: 222,
			Img: ``,
			ZH:  `关卡18`,
			En:  `Stage 18`,
			Ru:  `Этап 18`,
			DE:  `Stufe 18`,
			FR:  `Étape 18`,
		}, {
			Num: 223,
			Img: ``,
			ZH:  `关卡19`,
			En:  `Stage 19`,
			Ru:  `Этап 19`,
			DE:  `Stufe 19`,
			FR:  `Étape 19`,
		}, {
			Num: 224,
			Img: ``,
			ZH:  `关卡20`,
			En:  `Stage 20`,
			Ru:  `Этап 20`,
			DE:  `Stufe 20`,
			FR:  `Étape 20`,
		}, {
			Num: 225,
			Img: ``,
			ZH:  `关卡21`,
			En:  `Stage 21`,
			Ru:  `Этап 21`,
			DE:  `Stufe 21`,
			FR:  `Étape 21`,
		}, {
			Num: 226,
			Img: ``,
			ZH:  `关卡22`,
			En:  `Stage 22`,
			Ru:  `Этап 22`,
			DE:  `Stufe 22`,
			FR:  `Étape 22`,
		}, {
			Num: 227,
			Img: `列表暂无内容时提示`,
			ZH:  `暂无相关内容~`,
			En:  `No related content~`,
			Ru:  `Нет связанного контента`,
			DE:  `Kein dazugehöriger Inhalt`,
			FR:  `Aucun contenu associé`,
		}, {
			Num: 228,
			Img: ``,
			ZH:  `获取失败`,
			En:  `Failed to retrieve`,
			Ru:  `Не удалось получить`,
			DE:  `Abrufen fehlgeschlagen`,
			FR:  `Échec de la récupération`,
		}, {
			Num: 229,
			Img: ``,
			ZH:  `获取成功`,
			En:  `Successfully retrieved`,
			Ru:  `Успешно получено`,
			DE:  `Erfolgreich abgerufen`,
			FR:  `Récupéré avec succès`,
		}, {
			Num: 230,
			Img: ``,
			ZH:  `暂无相关神器`,
			En:  `No related artifacts`,
			Ru:  `Нет связанных артефактов`,
			DE:  `Keine dazugehörigen Artefakte`,
			FR:  `Aucun artefact associé`,
		}, {
			Num: 231,
			Img: ``,
			ZH:  `队列`,
			En:  `Queue`,
			Ru:  `Очередь`,
			DE:  `Warteschlange`,
			FR:  `File d'attente`,
		}, {
			Num: 232,
			Img: ``,
			ZH:  `反馈内容不能为空`,
			En:  `Feedback content cannot be empty`,
			Ru:  `Содержимое отзыва не может быть пустым`,
			DE:  `Die Rückmeldung darf nicht leer sein.`,
			FR:  `Le contenu du feedback ne peut pas être vide`,
		}, {
			Num: 233,
			Img: ``,
			ZH:  `反馈失败`,
			En:  `Feedback failed`,
			Ru:  `Отправка отзыва не удалась`,
			DE:  `Rückmeldung fehlgeschlagen`,
			FR:  `Échec du feedback`,
		}, {
			Num: 234,
			Img: ``,
			ZH:  `反馈成功`,
			En:  `Feedback successful`,
			Ru:  `Отзыв успешно отправлен`,
			DE:  `Rückmeldung erfolgreich`,
			FR:  `Feedback réussi`,
		}, {
			Num: 235,
			Img: ``,
			ZH:  `成功`,
			En:  `Success`,
			Ru:  `Успешно`,
			DE:  `Erfolgreich`,
			FR:  `Succès`,
		}, {
			Num: 236,
			Img: ``,
			ZH:  `操作成功`,
			En:  `Operation successful`,
			Ru:  `Операция успешна`,
			DE:  `Einsatz erfolgreich`,
			FR:  `Opération réussie`,
		}, {
			Num: 237,
			Img: ``,
			ZH:  `操作失败`,
			En:  `Operation failed`,
			Ru:  `Операция не удалась`,
			DE:  `Einsatz fehlgeschlagen`,
			FR:  `Opération échouée`,
		}, {
			Num: 240,
			Img: ``,
			ZH:  `未登录`,
			En:  `Not logged in`,
			Ru:  `Не выполнен вход`,
			DE:  `Nicht angemeldet`,
			FR:  `Déconnecté`,
		}, {
			Num: 241,
			Img: ``,
			ZH:  `退出后当前填写内容将不会被保存`,
			En:  `After logging out, the current content will not be saved`,
			Ru:  `После выхода, текущий контент не будет сохранен`,
			DE:  `Nachdem Abmelden wird der Inhalt nicht gespeichert`,
			FR:  `Après la déconnexion, le contenu actuel ne sera pas enregistré`,
		}, {
			Num: 242,
			Img: ``,
			ZH:  `请输入您的签名`,
			En:  `Please enter your signature`,
			Ru:  `Пожалуйста, введите вашу подпись`,
			DE:  `Bitte gib deine Signatur ein`,
			FR:  `Veuillez entrer votre signature`,
		}, {
			Num: 243,
			Img: ``,
			ZH:  `长按保存`,
			En:  `Long press to save`,
			Ru:  `Долгое нажатие для сохранения`,
			DE:  `Drücke lange zum speichern`,
			FR:  `Appuyez longuement pour enregistrer`,
		}, {
			Num: 244,
			Img: ``,
			ZH:  `生成图片中...`,
			En:  `Generating image...`,
			Ru:  `Создание изображения...`,
			DE:  `Bild wird generiert...`,
			FR:  `Génération de l'image...`,
		}, {
			Num: 245,
			Img: ``,
			ZH:  `已选内容将被清空，确定要清空吗？`,
			En:  `Selected content will be cleared. Are you sure you want to clear it?`,
			Ru:  `Выбранный контент будет очищен. Вы уверены что хотите продолжить?`,
			DE:  `Ausgewählter Inhalt wird geleert. Bist du dir sicher, dass du ihn leeren willst?`,
			FR:  `Le contenu sélectionné sera effacé. Êtes-vous sûr de vouloir le supprimer ?`,
		}, {
			Num: 246,
			Img: ``,
			ZH:  `请选择英雄`,
			En:  `Please select heroes`,
			Ru:  `Пожалуйста, выберите героев`,
			DE:  `Bitte wähle deine Helden aus`,
			FR:  `Veuillez sélectionner des héros`,
		}, {
			Num: 247,
			Img: ``,
			ZH:  `请选择英雄装备/神器`,
			En:  `Please select hero gears/artifacts`,
			Ru:  `Пожалуйста, выберите снаряжение/артефакты героя`,
			DE:  `Bitte wähle Helden Ausrüstung/Artefakte`,
			FR:  `Veuillez sélectionner les équipements/artefacts des héros`,
		}, {
			Num: 248,
			Img: ``,
			ZH:  `已选两个神器了`,
			En:  `Two artifacts have been selected`,
			Ru:  `Два артефакта были выбраны`,
			DE:  `Zwei Artefakte sind ausgewählt worden`,
			FR:  `Deux artefacts ont été sélectionnés`,
		}, {
			Num: 249,
			Img: ``,
			ZH:  `已选两个装备了`,
			En:  `Two pieces of equipment have been selected`,
			Ru:  `Две части снаряжения были выбраны`,
			DE:  `Zwei Ausrüstungsteile sind ausgewählt worden`,
			FR:  `Deux équipements ont été sélectionnés`,
		}, {
			Num: 250,
			Img: ``,
			ZH:  `未找到该关卡`,
			En:  `Level not found`,
			Ru:  `Уровень не найден`,
			DE:  `Level nicht gefunden`,
			FR:  `Niveau introuvable`,
		}, {
			Num: 251,
			Img: ``,
			ZH:  `目录不能为空`,
			En:  `Directory cannot be empty`,
			Ru:  `Каталог не может быть пустым`,
			DE:  `Das Verzeichnis darf nicht leer sein`,
			FR:  `Le dossier ne peut pas être vide`,
		}, {
			Num: 252,
			Img: ``,
			ZH:  `文件名不能为空`,
			En:  `File name cannot be empty`,
			Ru:  `Имя файла не может быть пустым`,
			DE:  `Der Dateiname darf nicht leer sein`,
			FR:  `Le nom du fichier ne peut pas être vide`,
		}, {
			Num: 253,
			Img: ``,
			ZH:  `文件类型错误`,
			En:  `Incorrect file type`,
			Ru:  `Неверный тип файла`,
			DE:  `Falscher Dateityp`,
			FR:  `Type de fichier incorrect`,
		}, {
			Num: 254,
			Img: ``,
			ZH:  `只能上传图片`,
			En:  `Only images can be uploaded`,
			Ru:  `Можно загружать только изображения`,
			DE:  `Es können nur Bilder hochgeladen werden`,
			FR:  `Seulement des images peuvent être téléchargées`,
		}, {
			Num: 255,
			Img: ``,
			ZH:  `获取上传组件失败:`,
			En:  `Failed to get the upload component:`,
			Ru:  `Не удалось загрузить компонент:`,
			DE:  `Die Upload-Komponente konnte nicht abgerufen werden:`,
			FR:  `Échec du téléchargement:`,
		}, {
			Num: 256,
			Img: ``,
			ZH:  `验证码错误`,
			En:  `Captcha error`,
			Ru:  `Ошибка в капче`,
			DE:  `Captcha fehler`,
			FR:  `Erreur de captcha`,
		}, {
			Num: 257,
			Img: ``,
			ZH:  `用户不存在`,
			En:  `User does not exist`,
			Ru:  `Пользователь не существует`,
			DE:  `Der Benutzer existiert nicht`,
			FR:  `L'utilisateur n'existe pas`,
		}, {
			Num: 258,
			Img: ``,
			ZH:  `修改密码失败`,
			En:  `Password change failed`,
			Ru:  `Не удалось измените пароль`,
			DE:  `Ändern des Passwortes fehlgeschlagen`,
			FR:  `Échec du changement de mot de passe`,
		}, {
			Num: 259,
			Img: ``,
			ZH:  `修改密码成功`,
			En:  `Password changed successfully`,
			Ru:  `Пароль успешно изменен`,
			DE:  `Ändern des Passwortes erfolgreich`,
			FR:  `Mot de passe modifié avec succès`,
		}, {
			Num: 260,
			Img: ``,
			ZH:  `未发送验证码`,
			En:  `Verification code not sent`,
			Ru:  `Код подтверждения не был отправлен`,
			DE:  `Der Verifikationscode wurde nicht gesendet`,
			FR:  `Code de vérification non envoyé`,
		}, {
			Num: 261,
			Img: ``,
			ZH:  `获取验证码过于频繁,请过段时间再获取`,
			En:  `Too frequent verification code requests, please try again later`,
			Ru:  `Слишком частые запросы кода подтверждения, пожалуйста, попробуйте позднее`,
			DE:  `Zu viele Verifikationscodeanfragen, versuch es bitte später nocheinmal`,
			FR:  `Demandes de code de vérification trop fréquentes, veuillez réessayer plus tard`,
		}, {
			Num: 262,
			Img: ``,
			ZH:  `验证码生成失败`,
			En:  `Failed to generate verification code`,
			Ru:  `Не удалось сгенерировать код подтверждения`,
			DE:  `Der Verifikationscode konnte nicht generiert werden.`,
			FR:  `Échec de la génération du code de vérification`,
		}, {
			Num: 263,
			Img: ``,
			ZH:  `您确定要离开此页面吗？`,
			En:  `Are you sure you want to leave this page?`,
			Ru:  `Вы уверенны, что хотите покинуть эту страницу?`,
			DE:  `Bist du dir sicher, dass du die Seite verlassen willst?`,
			FR:  `Êtes-vous sûr de vouloir quitter cette page ?`,
		}, {
			Num: 264,
			Img: ``,
			ZH:  `正在审核中`,
			En:  `Under review`,
			Ru:  `В процессе рассмотрения`,
			DE:  `Wird überprüft`,
			FR:  `En cours de vérification`,
		}, {
			Num: 265,
			Img: ``,
			ZH:  `审核失败,请重新上传`,
			En:  `Review failed, please upload again`,
			Ru:  `Проверка не пройдена, пожалуйста, загрузите снова`,
			DE:  `Überprüfung fehlgeschlagen, bitte lade erneut hoch`,
			FR:  `Échec de la vérification, veuillez télécharger à nouveau`,
		}, {
			Num: 266,
			Img: ``,
			ZH:  `请输入个性签名`,
			En:  `Please enter your personalized signature`,
			Ru:  `Пожалуйста, введите вашу личную подпись`,
			DE:  `Bitte trage eine personalisierte Signatur ein.`,
			FR:  `Veuillez entrer votre signature personnalisée`,
		}, {
			Num: 267,
			Img: ``,
			ZH:  `头像审核未通过`,
			En:  `Avatar review not passed`,
			Ru:  `Аватар не прошёл проверку`,
			DE:  `Avatar Überprüfung nicht bestanden`,
			FR:  `Avatar invalide`,
		}, {
			Num: 268,
			Img: ``,
			ZH:  `找不到对应攻略`,
			En:  `Corresponding guide not found`,
			Ru:  `Соответствующее руководство не найдено`,
			DE:  `Entsprechender Leitfaden nicht gefunden`,
			FR:  `Guide correspondant introuvable`,
		}, {
			Num: 269,
			Img: ``,
			ZH:  `举报失败`,
			En:  `Report failed`,
			Ru:  `Не удалось создать отчет`,
			DE:  `Meldung fehlgeschlagen`,
			FR:  `Échec du signalement`,
		}, {
			Num: 270,
			Img: ``,
			ZH:  `举报成功`,
			En:  `Report successful`,
			Ru:  `Отчет успешно создан`,
			DE:  `Meldung erfolgreich`,
			FR:  `Signalement réussi`,
		}, {
			Num: 271,
			Img: ``,
			ZH:  `无法重复举报`,
			En:  `Cannot report repeatedly`,
			Ru:  `Невозможно создать отчет повторно`,
			DE:  `Wiederholte Meldung nicht möglich`,
			FR:  `Impossible de signaler plusieurs fois`,
		}, {
			Num: 272,
			Img: ``,
			ZH:  `line up不能超过5个`,
			En:  `Lineup cannot exceed 5`,
			Ru:  `Расстановка не может превышать 5`,
			DE:  `Aufstellung kann nicht höher als 5 sein`,
			FR:  `La composition ne peut pas dépasser 5`,
		}, {
			Num: 273,
			Img: ``,
			ZH:  `找不到修改的攻略`,
			En:  `Corresponding guide not found to modify`,
			Ru:  `Соответствующее руководство не найдено для изменения`,
			DE:  `Entsprechender Leitfaden nicht gefunden, um ihn zu bearbeiten`,
			FR:  `Guide correspondant introuvable pour modification`,
		}, {
			Num: 274,
			Img: ``,
			ZH:  `点赞失败`,
			En:  `Like failed`,
			Ru:  `Не удалось поставить лайк`,
			DE:  `Like fehlgeschlagen`,
			FR:  `Échec du like`,
		}, {
			Num: 275,
			Img: ``,
			ZH:  `点赞成功`,
			En:  `Liked successfully`,
			Ru:  `Лайк успешно поставлен`,
			DE:  `Like erfolgreich`,
			FR:  `J'aime ajouté avec succès`,
		}, {
			Num: 276,
			Img: ``,
			ZH:  `取消点赞失败`,
			En:  `Unlike failed`,
			Ru:  `Не удалось отозвать лайк`,
			DE:  `Entfernen des Likes fehlgeschlagen`,
			FR:  `Échec du unlike`,
		}, {
			Num: 277,
			Img: ``,
			ZH:  `取消点赞成功`,
			En:  `Unliked successfully`,
			Ru:  `Лайк успешно отозван`,
			DE:  `Entfernen des Likes erfolgreich`,
			FR:  `J'aime retiré avec succès`,
		}, {
			Num: 278,
			Img: ``,
			ZH:  `收藏失败`,
			En:  `Collect failed`,
			Ru:  `Не удалось добавить в коллекцию`,
			DE:  `Zusammenführung fehlgeschlagen`,
			FR:  `Échec de la collecte`,
		}, {
			Num: 279,
			Img: ``,
			ZH:  `收藏成功`,
			En:  `Collected successfully`,
			Ru:  `Успешно добавлено в коллекцию`,
			DE:  `Lesezeichen gesetzt`,
			FR:  `Ajouté aux favoris`,
		}, {
			Num: 280,
			Img: ``,
			ZH:  `下架攻略失败`,
			En:  `Strategy taken down failed`,
			Ru:  `Не удалось убрать стратегию`,
			DE:  `Das Zurücknehmen der Strategie ist gescheitert`,
			FR:  `Échec de la suppression de la stratégie`,
		}, {
			Num: 281,
			Img: ``,
			ZH:  `下架攻略成功`,
			En:  `Strategy taken down successfully`,
			Ru:  `Стратегия успешно убрана`,
			DE:  `Das Zurücknehmen der Strategie war erfolgreich`,
			FR:  `Stratégie supprimée avec succès`,
		}, {
			Num: 282,
			Img: ``,
			ZH:  `取消收藏失败`,
			En:  `Uncollect failed`,
			Ru:  `Не удалось убрать из коллекции`,
			DE:  `Zusammenführung zurücknehmen fehlgeschlagen`,
			FR:  `Échec du déclassement`,
		}, {
			Num: 283,
			Img: ``,
			ZH:  `取消收藏成功`,
			En:  `Uncollected successfully`,
			Ru:  `Успешно удалено из коллекции`,
			DE:  `Lesezeichen entfernt`,
			FR:  `Retiré des favoris`,
		}, {
			Num: 284,
			Img: ``,
			ZH:  `您的帐户异地登陆或令牌失效`,
			En:  `Your account has logged in from a different location or the token is invalid`,
			Ru:  `В ваш аккаунт был выполнен вход из другого места или токен недействителен`,
			DE:  `Dieser Account hat sich von einem anderen Ort angemeldet oder der Token ist ungültig`,
			FR:  `Votre compte s'est connecté depuis un autre emplacement ou le jeton est invalide`,
		}, {
			Num: 285,
			Img: ``,
			ZH:  `授权已过期`,
			En:  `Authorization expired`,
			Ru:  `Время для авторизации вышло`,
			DE:  `Authorisierung abgelaufen`,
			FR:  `Autorisation expirée`,
		}, {
			Num: 286,
			Img: ``,
			ZH:  `权限不足`,
			En:  `Insufficient permissions`,
			Ru:  `Недостаточно прав`,
			DE:  `Unzureichende Berechtigungen`,
			FR:  `Permissions insuffisantes`,
		}, {
			Num: 287,
			Img: ``,
			ZH:  `找不到对应地图`,
			En:  `Corresponding map not found`,
			Ru:  `Соответствующая карта не найдена`,
			DE:  `Entsprechende Karte nicht gefunden`,
			FR:  `Carte correspondante introuvable`,
		}, {
			Num: 288,
			Img: ``,
			ZH:  `找不到对应英雄`,
			En:  `Corresponding hero not found`,
			Ru:  `Соответствующий герой не найден`,
			DE:  `Entsprechender Held nicht gefunden`,
			FR:  `Héros correspondant introuvable`,
		}, {
			Num: 289,
			Img: ``,
			ZH:  `找不到对应装备`,
			En:  `Corresponding equipment not found`,
			Ru:  `Соответствующее снаряжение не найдено`,
			DE:  `Entsprechende Ausrüstung nicht gefunden`,
			FR:  `Équipement correspondant introuvable`,
		}, {
			Num: 290,
			Img: ``,
			ZH:  `上下位英雄不能超过5个`,
			En:  `The number of upper and lower heroes cannot exceed 5`,
			Ru:  `Количество верхних и нижних героев не может превышать 5`,
			DE:  `Die Anzahl der oberen und unteren Helden darf nicht größer als 5 sein`,
			FR:  `Le nombre de héros supérieurs et inférieurs ne peut pas dépasser 5`,
		}, {
			Num: 291,
			Img: ``,
			ZH:  `上阵英雄不能超过10个`,
			En:  `The number of deployed heroes cannot exceed 10`,
			Ru:  `Количество выставленных героев не может превышать 10`,
			DE:  `Die Anzahl der eingesetzen Helden darf nicht größer als 10 sein`,
			FR:  `Le nombre de héros déployés ne peut pas dépasser 10`,
		}, {
			Num: 292,
			Img: ``,
			ZH:  `用户名请输入6-10位字符`,
			En:  `Username must be 6-10 characters`,
			Ru:  `Имя пользователя должно содержать 6-10 символов`,
			DE:  `Der Benutzername muss zwischen 6 und 10 Zeichen lang sein`,
			FR:  `Le nom d'utilisateur doit contenir entre 6 et 10 caractères`,
		}, {
			Num: 293,
			Img: ``,
			ZH:  `密码请输入8-20位字符`,
			En:  `Password must be 8-20 characters`,
			Ru:  `Пароль должен содержать 8-20 символов`,
			DE:  `Das Passwort muss zwischen 8 und 20 Zeichen lang sein.`,
			FR:  `Le mot de passe doit contenir entre 8 et 20 caractères`,
		}, {
			Num: 294,
			Img: ``,
			ZH:  `用户名不存在或者密码错误`,
			En:  `Username does not exist or password is incorrect`,
			Ru:  `Имя пользователя не существует или пароль введён неверно`,
			DE:  `Der Benutzername existiert nicht oder das Passwort ist falsch`,
			FR:  `Nom d'utilisateur ou mot de passe incorrect`,
		}, {
			Num: 295,
			Img: ``,
			ZH:  `用户被禁止登录`,
			En:  `User is banned from logging in`,
			Ru:  `Пользователю запрещен вход в систему`,
			DE:  `Das Anmelden für den Benutzer ist gesperrt`,
			FR:  `Utilisateur interdit de connexion`,
		}, {
			Num: 296,
			Img: ``,
			ZH:  `获取token失败`,
			En:  `Failed to get token`,
			Ru:  `Получить токен не удалось`,
			DE:  `Der Token konnte nicht abgerufen werden`,
			FR:  `Échec de l'obtention du jeton`,
		}, {
			Num: 297,
			Img: ``,
			ZH:  `登录成功`,
			En:  `Login successful`,
			Ru:  `Вход выполнен успешно`,
			DE:  `Anmelden erfolgreich`,
			FR:  `Connexion réussie`,
		}, {
			Num: 298,
			Img: ``,
			ZH:  `设置登录状态失败`,
			En:  `Failed to set login status`,
			Ru:  `Не удалось установить статус входа`,
			DE:  `Der Anmeldestatus konnte nicht geändert werden`,
			FR:  `Échec de la configuration du statut de connexion`,
		}, {
			Num: 299,
			Img: ``,
			ZH:  `jwt作废失败`,
			En:  `JWT invalidation failed`,
			Ru:  `Не удалось аннулировать JWT`,
			DE:  `JWT-Invalidierung fehlgeschlagen`,
			FR:  `Échec de l'invalidation JWT`,
		}, {
			Num: 300,
			Img: ``,
			ZH:  `注册失败!`,
			En:  `Registration failed!`,
			Ru:  `Регистрация не удалась`,
			DE:  `Registrierung fehlgeschlagen!`,
			FR:  `Échec de l'inscription !`,
		}, {
			Num: 301,
			Img: ``,
			ZH:  `未发送验证码或验证码已过期`,
			En:  `Verification code not sent or has expired`,
			Ru:  `Код подтверждения не был отправлен, или истек срок его действия`,
			DE:  `Verifikationscode nicht gesendet oder abgelaufen`,
			FR:  `Code de vérification non envoyé ou expiré`,
		}, {
			Num: 302,
			Img: ``,
			ZH:  `请重新发送验证码`,
			En:  `Please resend the verification code`,
			Ru:  `Пожалуйста, отправьте код подтверждения повторно`,
			DE:  `Bitte sende den Verifikationcode erneut`,
			FR:  `Veuillez renvoyer le code de vérification`,
		}, {
			Num: 303,
			Img: ``,
			ZH:  `注册成功`,
			En:  `Registration successful`,
			Ru:  `Регистрация успешна`,
			DE:  `Registrierung erfolgreich`,
			FR:  `Inscription réussie`,
		}, {
			Num: 304,
			Img: ``,
			ZH:  `获取验证码过于频繁,请在一分钟后获取!`,
			En:  `Too frequent verification code requests, please try again after one minute!`,
			Ru:  `Слишком частые запросы кода подтверждения, пожалуйста, попробуйте через 1 минуту!`,
			DE:  `Zu viele Verifikationscode anfragen, versuch es bitte in einer Minute nocheinmal!`,
			FR:  `Demandes de code de vérification trop fréquentes, veuillez réessayer dans une minute !`,
		}, {
			Num: 305,
			Img: ``,
			ZH:  `验证码生成失败`,
			En:  `Failed to generate verification code`,
			Ru:  `Не удалось сгенерировать код подтверждения`,
			DE:  `Der Verifikationscode konnte nicht generiert werden.`,
			FR:  `Échec de la génération du code de vérification`,
		}, {
			Num: 306,
			Img: ``,
			ZH:  `该游戏id已注册`,
			En:  `This game ID is already registered`,
			Ru:  `Этот игровой ID уже зарегистрирован`,
			DE:  `Diese Spiel-ID ist bereits registriert`,
			FR:  `Cet identifiant de jeu est déjà enregistré`,
		}, {
			Num: 307,
			Img: ``,
			ZH:  `验证码获取成功`,
			En:  `Verification code retrieved successfully`,
			Ru:  `Код подтверждения был успешно получен`,
			DE: `Der Verifikationscode wurde erfolgreich 
				DE: `Der Verifikationscode wurde erfolgreich 
			FR: `Code de vérification récupéré avec succès`,
		}, {
			Num: 308,
			Img: ``,
			ZH:  `请勾选隐私条款`,
			En:  `By clicking "Log in" or "Sign Up," you agree to our Terms of Service  and  Privacy Policy`,
			Ru:  `Нажимая кнопку "Войти" или "Зарегистрироваться", вы соглашаетесь с нашими Условиями предоставления услуг и Политикой конфиденциальности.`,
			DE:  `Indem Sie auf „Anmelden“ oder „Registrieren“ klicken, erklären Sie sich mit unseren Nutzungsbedingungen und unserer Datenschutzrichtlinie einverstanden`,
			FR:  `En cliquant sur "Se connecter" ou "S'inscrire", vous acceptez nos conditions d'utilisation et notre politique de confidentialité`,
		}, {
			Num: 309,
			Img: ``,
			ZH:  `请输入正确的用户名`,
			En:  `Please enter a valid username`,
			Ru:  `Пожалуйста, введите корректное имя пользователя`,
			DE:  `Bitte gib einen gültigen Benutzernamen ein`,
			FR:  `Veuillez entrer un nom d'utilisateur valide`,
		}, {
			Num: 310,
			Img: ``,
			ZH:  `请输入8位以上的密码`,
			En:  `Please enter a password longer than 8 characters`,
			Ru:  `Пожалуйста, введите пароль длиннее чем 8 символов`,
			DE:  `Bitte gib ein Passwort ein, dass länger als 8 Zeichen ist`,
			FR:  `Veuillez entrer un mot de passe contenant plus de 8 caractères`,
		}, {
			Num: 311,
			Img: ``,
			ZH:  `请勾选隐私条款`,
			En:  `Please check the privacy policy box`,
			Ru:  `Пожалуйста, проверьте поле политики конфиденциальности`,
			DE:  `Bitte kreuze das Feld "Datenschutz" an`,
			FR:  `Veuillez cocher la case de la politique de confidentialité`,
		}, {
			Num: 312,
			Img: ``,
			ZH:  `两次秘密输入不一致，请检查`,
			En:  `The two passwords do not match, please check`,
			Ru:  `Оба пароля должны совпадать, пожалуйста, проверьте еще раз`,
			DE:  `Die beiden Passwörter stimmen nicht überein, bitte überprüfen`,
			FR:  `Les deux mots de passe ne correspondent pas, veuillez vérifier`,
		}, {
			Num: 313,
			Img: ``,
			ZH:  `注册成功提示语`,
			En:  `Registration success prompt`,
			Ru:  `Сообщение об успешной регистрации`,
			DE:  `Registrierung erfolgreich`,
			FR:  `Message de réussite d'inscription`,
		}, {
			Num: 314,
			Img: ``,
			ZH:  `暂无任何消息`,
			En:  `No messages`,
			Ru:  `Нет сообщений`,
			DE:  `Keine Nachrichten`,
			FR:  `Aucun message`,
		}, {
			Num: 315,
			Img: ``,
			ZH:  `已选择`,
			En:  `Selected`,
			Ru:  `Выбрано`,
			DE:  `Ausgewählt`,
			FR:  `Sélectionné`,
		}, {
			Num: 316,
			Img: ``,
			ZH:  `选择地图尺寸`,
			En:  `Select map size`,
			Ru:  `Выберите размер карты`,
			DE:  `Wähle die Kartengröße aus`,
			FR:  `Sélectionnez la taille de la carte`,
		}, {
			Num: 317,
			Img: ``,
			ZH:  `请上传图片证据`,
			En:  `Please upload image evidence`,
			Ru:  `Пожалуйста, загрузите изображение доказательства`,
			DE:  `Bitte lade einen Bildbeweis hoch`,
			FR:  `Veuillez télécharger des preuves en image`,
		}, {
			Num: 318,
			Img: ``,
			ZH:  `退出`,
			En:  `Log out`,
			Ru:  `Выйти`,
			DE:  `Abmelden`,
			FR:  `Se déconnecter`,
		}, {
			Num: 319,
			Img: ``,
			ZH:  `通用`,
			En:  `Generic`,
			Ru:  `Универсальн`,
			DE:  `Allgemein`,
			FR:  `Général`,
		}, {
			Num: 320,
			Img: ``,
			ZH:  `专属`,
			En:  `Exclusive`,
			Ru:  `Эксклюзивн`,
			DE:  `Exklusiv`,
			FR:  `Exclusif`,
		}, {
			Num: 321,
			Img: ``,
			ZH:  `确定下架该内容吗？`,
			En:  `Are you sure you want to unlist this content?`,
			Ru:  `Вы уверены, что хотите снять этот материал с публикации?`,
			DE:  `Bist du sicher, dass du diesen Inhalt entfernen möchtest?`,
			FR:  `Êtes-vous sûr de vouloir dépublier ce contenu ?`,
		}, {
			Num: 322,
			Img: ``,
			ZH:  `请构建地图`,
			En:  `Please build the map`,
			Ru:  `Пожалуйста, создайте карту`,
			DE:  `Bitte baue die Karte`,
			FR:  `Veuillez créer la carte`,
		}, {
			Num: 323,
			Img: ``,
			ZH:  `保存成功`,
			En:  `Saved successfully`,
			Ru:  `Успешно сохранено`,
			DE:  `Erfolgreich gespeichert`,
			FR:  `Enregistré avec succès`,
		}, {
			Num: 324,
			Img: ``,
			ZH:  `创建成功`,
			En:  `Created successfully`,
			Ru:  `Успешно создано`,
			DE:  `Erfolgreich erstellt`,
			FR:  `Créé avec succès`,
		}, {
			Num: 325,
			Img: ``,
			ZH:  `编辑成功`,
			En:  `Edited successfully`,
			Ru:  `Успешно отредактировано`,
			DE:  `Erfolgreich geändert`,
			FR:  `Modifié avec succès`,
		}, {
			Num: 326,
			Img: ``,
			ZH:  `获取上传组件失败,请重试`,
			En:  `Failed to retrieve the uploaded component. Please try again`,
			Ru:  `Не удалось загрузить компонент. Попробуйте снова`,
			DE:  `Die hochgeladene Komponente konnte nicht abgerufen werden. Bitte versuche es erneut`,
			FR:  `Échec du téléchargement. Veuillez réessayer`,
		}, {
			Num: 327,
			Img: ``,
			ZH:  `最多选10个英雄`,
			En:  `You can select up to 10 heroes`,
			Ru:  `Вы можете выбрать до 10 героев`,
			DE:  `Du kannst bis zu 10 Helden auswählen`,
			FR:  `Vous pouvez sélectionner jusqu'à 10 héros`,
		}, {
			Num: 328,
			Img: ``,
			ZH:  `最多可以选择5个`,
			En:  `You can select up to 5`,
			Ru:  `Вы можете выбрать до 5`,
			DE:  `Du kannst bis zu 5 auswählen`,
			FR:  `Vous pouvez en sélectionner jusqu'à 5`,
		}, {
			Num: 329,
			Img: ``,
			ZH:  `发布成功`,
			En:  `Published successfully`,
			Ru:  `Успешно опубликовано`,
			DE:  `Erfolgreich veröffentlicht`,
			FR:  `Publié avec succès`,
		}, {
			Num: 330,
			Img: ``,
			ZH:  `请输入关键词`,
			En:  `Please enter keywords`,
			Ru:  `Введите ключевые слова`,
			DE:  `Bitte Schlüsselwörter eingeben`,
			FR:  `Veuillez entrer les mots-clés`,
		}, {
			Num: 331,
			Img: ``,
			ZH:  `请输入签名`,
			En:  `Please enter your signature`,
			Ru:  `Введите вашу подпись`,
			DE:  `Bitte gib deine Unterschrift ein`,
			FR:  `Veuillez entrer votre signature`,
		}, {
			Num: 332,
			Img: ``,
			ZH:  `继续填写`,
			En:  `Continue editing`,
			Ru:  `Продолжить редактирование`,
			DE:  `Bearbeitung fortsetzen`,
			FR:  `Poursuivre la modification`,
		}, {
			Num: 333,
			Img: ``,
			ZH:  `联系客服`,
			En:  `Contact customer service`,
			Ru:  `Свяжитесь со службой поддержки`,
			DE:  `Kontakt zum Kundendienst`,
			FR:  `Contacter le service client`,
		}, {
			Num: 334,
			Img: ``,
			ZH:  `制作我自己的攻略`,
			En:  `Create my own guide`,
			Ru:  `Создать собственное руководство`,
			DE:  `Meinen eigenen Leitfaden erstellen`,
			FR:  `Créer mon propre guide`,
		}, {
			Num: 335,
			Img: ``,
			ZH:  `举报`,
			En:  `Report`,
			Ru:  `Сообщить о проблеме`,
			DE:  `Melden`,
			FR:  `Signaler`,
		}, {
			Num: 336,
			Img: ``,
			ZH:  `选择英雄配装`,
			En:  `Select Hero Gear Loadout`,
			Ru:  `Выберите Комплект снаряжения героя`,
			DE:  `Heldenausrüstung auswählen`,
			FR:  `Sélectionnez Configuration d'équipement de héros`,
		}, {
			Num: 337,
			Img: ``,
			ZH:  `登录过于频繁,请稍后再登录`,
			En:  `Login is too frequent, please try again later`,
			Ru:  `Слишком частые попытки входа, пожалуйста, попробуйте позже`,
			DE:  `Zu häufige Anmeldungen, bitte versuchen Sie es später erneut`,
			FR:  `Connexions trop fréquentes, veuillez réessayer plus tard`,
		}, {
			Num: 338,
			Img: ``,
			ZH:  `头像审核通过`,
			En:  `Avatar approved`,
			Ru:  `Аватар одобрен`,
			DE:  `Avatar genehmigt`,
			FR:  `Avatar approuvé`,
		}, {
			Num: 339,
			Img: ``,
			ZH:  `用户名已注册`,
			En:  `The username is already registered.`,
			Ru:  `Имя пользователя уже зарегистрировано.`,
			DE:  `Der Benutzername ist bereits registriert.`,
			FR:  `Le nom d'utilisateur est déjà enregistré.`,
		}, {
			Num: 340,
			Img: ``,
			ZH:  ``,
			En:  `Spawn`,
			Ru:  `Точка появления`,
			DE:  `Spawnpunkt`,
			FR:  `Point d'apparition`,
		}, {
			Num: 341,
			Img: ``,
			ZH:  ``,
			En:  `Walls`,
			Ru:  `Стены`,
			DE:  `Wände`,
			FR:  `Murs`,
		}, {
			Num: 342,
			Img: ``,
			ZH:  ``,
			En:  `Elevate`,
			Ru:  `Поднять`,
			DE:  `Erhöhen`,
			FR:  `Élever`,
		}, {
			Num: 343,
			Img: ``,
			ZH:  ``,
			En:  `Ground`,
			Ru:  `Земля`,
			DE:  `Boden`,
			FR:  `Sol`,
		}, {
			Num: 344,
			Img: ``,
			ZH:  ``,
			En:  `Core`,
			Ru:  `Ядро`,
			DE:  `Kern`,
			FR:  `Noyau`,
		}, {
			Num: 345,
			Img: ``,
			ZH:  ``,
			En:  `Route-right`,
			Ru:  `Маршрут направо`,
			DE:  `Route-rechts`,
			FR:  `Route à droite`,
		}, {
			Num: 346,
			Img: ``,
			ZH:  ``,
			En:  `Route-left`,
			Ru:  `Маршрут налево`,
			DE:  `Route-links`,
			FR:  `Route à gauche`,
		}, {
			Num: 347,
			Img: ``,
			ZH:  ``,
			En:  `Route-top`,
			Ru:  `Маршрут вверх`,
			DE:  `Route-oben`,
			FR:  `Route en haut`,
		}, {
			Num: 348,
			Img: ``,
			ZH:  ``,
			En:  `Route-bottom`,
			Ru:  `Маршрут вниз`,
			DE:  `Route-unten`,
			FR:  `Route en bas`,
		}, {
			Num: 349,
			Img: ``,
			ZH:  ``,
			En:  `Route-left-top`,
			Ru:  `Маршрут налево-вверх`,
			DE:  `Route-links-oben`,
			FR:  `Route en haut à gauche`,
		}, {
			Num: 350,
			Img: ``,
			ZH:  ``,
			En:  `Route-right-top`,
			Ru:  `Маршрут направо-вверх`,
			DE:  `Route-rechts-oben`,
			FR:  `Route en haut à droite`,
		}, {
			Num: 351,
			Img: ``,
			ZH:  ``,
			En:  `Route-left-bottom`,
			Ru:  `Маршрут налево-вниз`,
			DE:  `Route-links-unten`,
			FR:  `Route en bas à gauche`,
		}, {
			Num: 352,
			Img: ``,
			ZH:  ``,
			En:  `Route-right-bottom`,
			Ru:  `Маршрут направо-вниз`,
			DE:  `Route-rechts-unten`,
			FR:  `Route en bas à droite`,
		}, {
			Num: 353,
			Img: ``,
			ZH:  ``,
			En:  `Mobs`,
			Ru:  `Мобы`,
			DE:  `Gegner`,
			FR:  `Ennemis`,
		}, {
			Num: 354,
			Img: ``,
			ZH:  ``,
			En:  `Fence`,
			Ru:  `Забор`,
			DE:  `Zaun`,
			FR:  `Clôture`,
		}, {
			Num: 355,
			Img: ``,
			ZH:  `By clicking "Log in" or "Sign Up," you agree to our Terms of Service and Privacy Policy`,
			En:  `By clicking "Log in" or "Sign Up," you agree to our Terms of <a href="https://www.watcherofrealms.com/?config=page_term" target="_blank">Service</a> and <a href="https://www.watcherofrealms.com/?config=page_policy" target="_blank">Privacy Policy</a>`,
			Ru:  `Нажимая «Войти» или «Зарегистрироваться», вы соглашаетесь с нашими <a href="https://www.watcherofrealms.com/?config=page_term" target="_blank">Условиями использования</a> и <a href="https://www.watcherofrealms.com/?config=page_policy" target="_blank">Политикой конфиденциальности</a>.`,
			DE:  `Durch Klicken auf „Anmelden“ oder „Registrieren“ stimmen Sie unseren <a href="https://www.watcherofrealms.com/?config=page_term" target="_blank">Nutzungsbedingungen</a> und unserer <a href="https://www.watcherofrealms.com/?config=page_policy" target="_blank">Datenschutzrichtlinie</a> zu.`,
			FR:  `En cliquant sur « Se connecter » ou « S'inscrire », vous acceptez nos <a href="https://www.watcherofrealms.com/?config=page_term" target="_blank">Conditions d'utilisation</a> et notre <a href="https://www.watcherofrealms.com/?config=page_policy" target="_blank">Politique de confidentialité</a>.`,
		}, {
			Num: 356,
			Img: ``,
			ZH:  `攻略已下架`,
			En:  `The guide has been unlisted`,
			Ru:  `Прохождение было скрыто из списка`,
			DE:  `Die Anleitung wurde aus der Liste entfernt`,
			FR:  `Le guide a été retiré de la liste`,
		}, {
			Num: 357,
			Img: ``,
			ZH:  ``,
			En:  `Line Up  1`,
			Ru:  `Состав  1`,
			DE:  ``,
			FR:  `Équipe 1`,
		}, {
			Num: 358,
			Img: ``,
			ZH:  `暂无相关内容`,
			En:  `No results`,
			Ru:  `Нет результатов`,
			DE:  `Keine Ergebnisse`,
			FR:  `Pas de résultat`,
		}, {
			Num: 359,
			Img: ``,
			ZH:  `生成中`,
			En:  `Generating...`,
			Ru:  `Генерация...`,
			DE:  `Wird generiert...`,
			FR:  `Génération en cours...`,
		}, {
			Num: 360,
			Img: ``,
			ZH:  `长按保存`,
			En:  `Long press to save`,
			Ru:  `Долгое нажатие для сохранения`,
			DE:  `Lang drücken, um zu speichern`,
			FR:  `Appuyez longuement pour enregistrer`,
		}, {
			Num: 361,
			Img: ``,
			ZH:  `请输入您的签名`,
			En:  `Please enter your signature`,
			Ru:  `Пожалуйста, введите вашу подпись`,
			DE:  `Bitte geben Sie Ihre Signatur ein`,
			FR:  `Veuillez entrer votre signature`,
		}, {
			Num: 362,
			Img: ``,
			ZH:  `编辑我的个性签名`,
			En:  `Edit my personal signature`,
			Ru:  `Редактировать мою персональную подпись`,
			DE:  `Meine persönliche Signatur bearbeiten`,
			FR:  `Modifier ma signature personnelle`,
		}, {
			Num: 363,
			Img: ``,
			ZH:  `已上阵`,
			En:  `Deployed`,
			Ru:  `Размещено`,
			DE:  `Eingesetzt`,
			FR:  `Déployé`,
		}, {
			Num: 364,
			Img: ``,
			ZH:  `请输入正确的区服ID`,
			En:  `Please enter the correct server ID`,
			Ru:  `Пожалуйста, введите правильный ID сервера`,
			DE:  `Bitte geben Sie die korrekte Server-ID ein`,
			FR:  `Veuillez entrer l'ID du serveur correct`,
		}, {
			Num: 365,
			Img: ``,
			ZH:  `不可选`,
			En:  `Unavailable`,
			Ru:  `Недоступно`,
			DE:  `Nicht verfügbar`,
			FR:  `Indisponible`,
		},
	}
)

	)
	/*用户名*/
	Num1 = 1
	/*请输入用户名*/
	Num2 = 2
	/*密码*/
	Num3 = 3
	/*请输入登录密码*/
	Num4 = 4
	/*请输入邮箱验证码*/
	Num5 = 5
	/*发送验证码*/
	Num6 = 6
	/*请前往游戏内邮箱获取验证码*/
	Num7 = 7
	/*登录*/
	Num8 = 8
	/*注册账户*/
	Num9 = 9
	/*忘记密码*/
	Num10 = 10
	/*登录*/
	Num11 = 11
	/*注册*/
	Num12 = 12
	/*用户名*/
	Num13 = 13
	/*请输入纯字母或数字 不超过10个字符*/
	Num14 = 14
	/*游戏ID*/
	Num15 = 15
	/*请输入游戏ID*/
	Num16 = 16
	/*区服ID*/
	Num17 = 17
	/*请输入区服ID*/
	Num18 = 18
	/*请输入邮箱验证码*/
	Num19 = 19
	/*发送验证码*/
	Num20 = 20
	/*请前往游戏内邮箱获取验证码*/
	Num21 = 21
	/*密码*/
	Num22 = 22
	/*请输入登录密码*/
	Num23 = 23
	/*密码确认*/
	Num24 = 24
	/*请再次输入登录密码*/
	Num25 = 25
	/*注册并登录*/
	Num26 = 26
	/*修改密码*/
	Num27 = 27
	/*新密码*/
	Num28 = 28
	/*请输入新的登录密码*/
	Num29 = 29
	/*密码确认*/
	Num30 = 30
	/*游戏ID*/
	Num31 = 31
	/*请输入游戏ID*/
	Num32 = 32
	/*区服ID*/
	Num33 = 33
	/*请输入区服ID*/
	Num34 = 34
	/*请输入邮箱验证码*/
	Num35 = 35
	/*发送验证码*/
	Num36 = 36
	/*请前往游戏内邮箱获取验证码*/
	Num37 = 37
	/*修改密码*/
	Num38 = 38
	/*提示*/
	Num39 = 39
	/*当前登录账户因违反社区规则*/
	Num40 = 40
	/*账户已经封停至xxx年xx月xx日xx:xx*/
	Num41 = 41
	/*密码修改成功
	  请使用新密码进行登录*/
	Num42 = 42
	/*登录*/
	Num43 = 43
	/*请先登录 再进行操作*/
	Num44 = 44
	/*用户名*/
	Num45 = 45
	/*请输入用户名*/
	Num46 = 46
	/*密码*/
	Num47 = 47
	/*请输入登录密码*/
	Num48 = 48
	/*请输入邮箱验证码*/
	Num49 = 49
	/*发送验证码*/
	Num50 = 50
	/*请前往游戏内邮箱获取验证码*/
	Num51 = 51
	/*注册账户*/
	Num52 = 52
	/*攻略制作器*/
	Num53 = 53
	/*我的书签*/
	Num54 = 54
	/*搜索*/
	Num55 = 55
	/*筛选*/
	Num56 = 56
	/*全部重置*/
	Num57 = 57
	/*选择游戏关卡*/
	Num58 = 58
	/*显示所有早期游戏内容 推荐攻略内容*/
	Num59 = 59
	/*最新*/
	Num60 = 60
	/*最热*/
	Num61 = 61
	/*点赞*/
	Num62 = 62
	/*收藏*/
	Num63 = 63
	/*首页*/
	Num64 = 64
	/*编辑器*/
	Num65 = 65
	/*书签*/
	Num66 = 66
	/*我的*/
	Num67 = 67
	/*关卡选择*/
	Num68 = 68
	/*取消*/
	Num69 = 69
	/*确定*/
	Num70 = 70
	/*我的书签*/
	Num71 = 71
	/*攻略已下架*/
	Num72 = 72
	/*消息列表*/
	Num73 = 73
	/*消息*/
	Num74 = 74
	/*WOR指南生成器*/
	Num75 = 75
	/*标题*/
	Num76 = 76
	/*每篇好的攻略都需要一个伟大的标题（20个字）*/
	Num77 = 77
	/*描述*/
	Num78 = 78
	/*简单描述你的好指南（200个字）
	  例如: 1.主c英雄的战力区间
	  2.阵营小队激活效果
	  3.切line up的时机*/
	Num79 = 79
	/*排序*/
	Num80 = 80
	/*全部重置*/
	Num81 = 81
	/*请配置你的英雄阵容，英雄分为可被上下位替换和不可替换英雄*/
	Num82 = 82
	/*位置*/
	Num83 = 83
	/*你可以选择已经预制好的地图，也可以自定义编辑关卡地图,然后按地图上的空格，选择你的英雄和他们的定位；*/
	Num84 = 84
	/*选择您的地图*/
	Num85 = 85
	/*自定义地图*/
	Num86 = 86
	/*1.在选项卡之间切换，为不同的战斗阶段设置英雄
	  2.点击色块，选择上阵英雄和英雄方向*/
	Num87 = 87
	/*line up1 的描述内容*/
	Num88 = 88
	/*配置英雄装备*/
	Num89 = 89
	/*简述你的配装思路（100个字符）*/
	Num90 = 90
	/*英雄*/
	Num91 = 91
	/*装备效果*/
	Num92 = 92
	/*神器*/
	Num93 = 93
	/*点击英雄头像修改配置*/
	Num94 = 94
	/*专属神器*/
	Num95 = 95
	/*通用神器*/
	Num96 = 96
	/*(选填)请简单描述装备套装，如主副属性、强化目标等(50个字符)*/
	Num97 = 97
	/*添加英雄*/
	Num98 = 98
	/*攻略视频*/
	Num99 = 99
	/*(选填)攻略视频链接*/
	Num100 = 100
	/*预览*/
	Num101 = 101
	/*发布*/
	Num102 = 102
	/*保存为长图*/
	Num103 = 103
	/*存为草稿*/
	Num104 = 104
	/*上阵英雄选择*/
	Num105 = 105
	/*选择上阵英雄（必选）*/
	Num106 = 106
	/*选择英雄方向（必选）*/
	Num107 = 107
	/*向上*/
	Num108 = 108
	/*向下*/
	Num109 = 109
	/*向左*/
	Num110 = 110
	/*向右*/
	Num111 = 111
	/*上阵顺序（必选）*/
	Num112 = 112
	/*请选择上阵英雄序号*/
	Num113 = 113
	/*确定*/
	Num114 = 114
	/*上阵英雄配置*/
	Num115 = 115
	/*选择配置英雄（必选）*/
	Num116 = 116
	/*已配置*/
	Num117 = 117
	/*选择英雄装备效果（非必选）*/
	Num118 = 118
	/*武器胸甲*/
	Num119 = 119
	/*首饰*/
	Num120 = 120
	/*选择英雄神器（非必选）*/
	Num121 = 121
	/*先选择的为首选神器后选择的为次选神器*/
	Num122 = 122
	/*法师*/
	Num123 = 123
	/*射手*/
	Num124 = 124
	/*守护者*/
	Num125 = 125
	/*医师*/
	Num126 = 126
	/*战士*/
	Num127 = 127
	/*3星*/
	Num128 = 128
	/*4星*/
	Num129 = 129
	/*5星*/
	Num130 = 130
	/*保存*/
	Num131 = 131
	/*提示*/
	Num132 = 132
	/*当前内容未保存*/
	Num133 = 133
	/*返回*/
	Num134 = 134
	/*保存草稿箱*/
	Num135 = 135
	/*选择英雄*/
	Num136 = 136
	/*保存*/
	Num137 = 137
	/*英雄池*/
	Num138 = 138
	/*点击英雄头像来上阵该英雄*/
	Num139 = 139
	/*不可替代英雄*/
	Num140 = 140
	/*法师*/
	Num141 = 141
	/*射手*/
	Num142 = 142
	/*守护者*/
	Num143 = 143
	/*医师*/
	Num144 = 144
	/*战士*/
	Num145 = 145
	/*已选择*/
	Num146 = 146
	/*可被替代英雄*/
	Num147 = 147
	/*提示*/
	Num148 = 148
	/*是否保存当前阵容*/
	Num149 = 149
	/*返回*/
	Num150 = 150
	/*保存*/
	Num151 = 151
	/*关卡地图选择*/
	Num152 = 152
	/*选择关卡*/
	Num153 = 153
	/*关卡地图选择*/
	Num154 = 154
	/*选择关卡*/
	Num155 = 155
	/*请选择上下位英雄*/
	Num156 = 156
	/*英雄池*/
	Num157 = 157
	/*点击英雄头像来上阵该英雄*/
	Num158 = 158
	/*可被替代英雄*/
	Num159 = 159
	/*已选择*/
	Num160 = 160
	/*保存*/
	Num161 = 161
	/*WOR地图编辑器*/
	Num162 = 162
	/*选择绘制的地图关卡*/
	Num163 = 163
	/*请选择地图的长度*/
	Num164 = 164
	/*请选择地图的宽度*/
	Num165 = 165
	/*保存*/
	Num166 = 166
	/*确定*/
	Num167 = 167
	/*取消*/
	Num168 = 168
	/*视频链接*/
	Num169 = 169
	/*地图位置*/
	Num170 = 170
	/*英雄装备攻略*/
	Num171 = 171
	/*制作我自己的攻略*/
	Num172 = 172
	/*举报反馈*/
	Num173 = 173
	/*作者名*/
	Num174 = 174
	/*攻略标题*/
	Num175 = 175
	/*举报原因（必填）*/
	Num176 = 176
	/*请输入原因*/
	Num177 = 177
	/*图片证据（最多上传9张）*/
	Num178 = 178
	/*提交*/
	Num179 = 179
	/*提示*/
	Num180 = 180
	/*即将离开本页面，请注意账号安全*/
	Num181 = 181
	/*继续访问*/
	Num182 = 182
	/*提示*/
	Num183 = 183
	/*提交成功，我们将尽快处理*/
	Num184 = 184
	/*我的攻略*/
	Num185 = 185
	/*已下架*/
	Num186 = 186
	/*草稿箱*/
	Num187 = 187
	/*下架*/
	Num188 = 188
	/*删除*/
	Num189 = 189
	/*提示*/
	Num190 = 190
	/*确定删除该内容吗？*/
	Num191 = 191
	/*确定*/
	Num192 = 192
	/*取消*/
	Num193 = 193
	/*建议反馈*/
	Num194 = 194
	/*编辑*/
	Num195 = 195
	/*建议（必填）*/
	Num196 = 196
	/*请输入内容*/
	Num197 = 197
	/*上传图片（最多上传9张）*/
	Num198 = 198
	/*装备副本I*/
	Num199 = 199
	/*装备副本II*/
	Num200 = 200
	/*装备副本III*/
	Num201 = 201
	/*神器材料副本*/
	Num202 = 202
	/*装备地下城I*/
	Num203 = 203
	/*装备地下城II*/
	Num204 = 204
	/*关卡1*/
	Num205 = 205
	/*关卡2*/
	Num206 = 206
	/*关卡3*/
	Num207 = 207
	/*关卡4*/
	Num208 = 208
	/*关卡5*/
	Num209 = 209
	/*关卡6*/
	Num210 = 210
	/*关卡7*/
	Num211 = 211
	/*关卡8*/
	Num212 = 212
	/*关卡9*/
	Num213 = 213
	/*关卡10*/
	Num214 = 214
	/*关卡11*/
	Num215 = 215
	/*关卡12*/
	Num216 = 216
	/*关卡13*/
	Num217 = 217
	/*关卡14*/
	Num218 = 218
	/*关卡15*/
	Num219 = 219
	/*关卡16*/
	Num220 = 220
	/*关卡17*/
	Num221 = 221
	/*关卡18*/
	Num222 = 222
	/*关卡19*/
	Num223 = 223
	/*关卡20*/
	Num224 = 224
	/*关卡21*/
	Num225 = 225
	/*关卡22*/
	Num226 = 226
	/*暂无相关内容~*/
	Num227 = 227
	/*获取失败*/
	Num228 = 228
	/*获取成功*/
	Num229 = 229
	/*暂无相关神器*/
	Num230 = 230
	/*队列*/
	Num231 = 231
	/*反馈内容不能为空*/
	Num232 = 232
	/*反馈失败*/
	Num233 = 233
	/*反馈成功*/
	Num234 = 234
	/*成功*/
	Num235 = 235
	/*操作成功*/
	Num236 = 236
	/*操作失败*/
	Num237 = 237
	/*未登录*/
	Num240 = 240
	/*退出后当前填写内容将不会被保存*/
	Num241 = 241
	/*请输入您的签名*/
	Num242 = 242
	/*长按保存*/
	Num243 = 243
	/*生成图片中...*/
	Num244 = 244
	/*已选内容将被清空，确定要清空吗？*/
	Num245 = 245
	/*请选择英雄*/
	Num246 = 246
	/*请选择英雄装备/神器*/
	Num247 = 247
	/*已选两个神器了*/
	Num248 = 248
	/*已选两个装备了*/
	Num249 = 249
	/*未找到该关卡*/
	Num250 = 250
	/*目录不能为空*/
	Num251 = 251
	/*文件名不能为空*/
	Num252 = 252
	/*文件类型错误*/
	Num253 = 253
	/*只能上传图片*/
	Num254 = 254
	/*获取上传组件失败:*/
	Num255 = 255
	/*验证码错误*/
	Num256 = 256
	/*用户不存在*/
	Num257 = 257
	/*修改密码失败*/
	Num258 = 258
	/*修改密码成功*/
	Num259 = 259
	/*未发送验证码*/
	Num260 = 260
	/*获取验证码过于频繁,请过段时间再获取*/
	Num261 = 261
	/*验证码生成失败*/
	Num262 = 262
	/*您确定要离开此页面吗？*/
	Num263 = 263
	/*正在审核中*/
	Num264 = 264
	/*审核失败,请重新上传*/
	Num265 = 265
	/*请输入个性签名*/
	Num266 = 266
	/*头像审核未通过*/
	Num267 = 267
	/*找不到对应攻略*/
	Num268 = 268
	/*举报失败*/
	Num269 = 269
	/*举报成功*/
	Num270 = 270
	/*无法重复举报*/
	Num271 = 271
	/*line up不能超过5个*/
	Num272 = 272
	/*找不到修改的攻略*/
	Num273 = 273
	/*点赞失败*/
	Num274 = 274
	/*点赞成功*/
	Num275 = 275
	/*取消点赞失败*/
	Num276 = 276
	/*取消点赞成功*/
	Num277 = 277
	/*收藏失败*/
	Num278 = 278
	/*收藏成功*/
	Num279 = 279
	/*下架攻略失败*/
	Num280 = 280
	/*下架攻略成功*/
	Num281 = 281
	/*取消收藏失败*/
	Num282 = 282
	/*取消收藏成功*/
	Num283 = 283
	/*您的帐户异地登陆或令牌失效*/
	Num284 = 284
	/*授权已过期*/
	Num285 = 285
	/*权限不足*/
	Num286 = 286
	/*找不到对应地图*/
	Num287 = 287
	/*找不到对应英雄*/
	Num288 = 288
	/*找不到对应装备*/
	Num289 = 289
	/*上下位英雄不能超过5个*/
	Num290 = 290
	/*上阵英雄不能超过10个*/
	Num291 = 291
	/*用户名请输入6-10位字符*/
	Num292 = 292
	/*密码请输入8-20位字符*/
	Num293 = 293
	/*用户名不存在或者密码错误*/
	Num294 = 294
	/*用户被禁止登录*/
	Num295 = 295
	/*获取token失败*/
	Num296 = 296
	/*登录成功*/
	Num297 = 297
	/*设置登录状态失败*/
	Num298 = 298
	/*jwt作废失败*/
	Num299 = 299
	/*注册失败!*/
	Num300 = 300
	/*未发送验证码或验证码已过期*/
	Num301 = 301
	/*请重新发送验证码*/
	Num302 = 302
	/*注册成功*/
	Num303 = 303
	/*获取验证码过于频繁,请在一分钟后获取!*/
	Num304 = 304
	/*验证码生成失败*/
	Num305 = 305
	/*该游戏id已注册*/
	Num306 = 306
	/*验证码获取成功*/
	Num307 = 307
	/*请勾选隐私条款*/
	Num308 = 308
	/*请输入正确的用户名*/
	Num309 = 309
	/*请输入8位以上的密码*/
	Num310 = 310
	/*请勾选隐私条款*/
	Num311 = 311
	/*两次秘密输入不一致，请检查*/
	Num312 = 312
	/*注册成功提示语*/
	Num313 = 313
	/*暂无任何消息*/
	Num314 = 314
	/*已选择*/
	Num315 = 315
	/*选择地图尺寸*/
	Num316 = 316
	/*请上传图片证据*/
	Num317 = 317
	/*退出*/
	Num318 = 318
	/*通用*/
	Num319 = 319
	/*专属*/
	Num320 = 320
	/*确定下架该内容吗？*/
	Num321 = 321
	/*请构建地图*/
	Num322 = 322
	/*保存成功*/
	Num323 = 323
	/*创建成功*/
	Num324 = 324
	/*编辑成功*/
	Num325 = 325
	/*获取上传组件失败,请重试*/
	Num326 = 326
	/*最多选10个英雄*/
	Num327 = 327
	/*最多可以选择5个*/
	Num328 = 328
	/*发布成功*/
	Num329 = 329
	/*请输入关键词*/
	Num330 = 330
	/*请输入签名*/
	Num331 = 331
	/*继续填写*/
	Num332 = 332
	/*联系客服*/
	Num333 = 333
	/*制作我自己的攻略*/
	Num334 = 334
	/*举报*/
	Num335 = 335
	/*选择英雄配装*/
	Num336 = 336
	/*登录过于频繁,请稍后再登录*/
	Num337 = 337
	/*头像审核通过*/
	Num338 = 338
	/*用户名已注册*/
	Num339 = 339
	/**/
	Num340 = 340
	/**/
	Num341 = 341
	/**/
	Num342 = 342
	/**/
	Num343 = 343
	/**/
	Num344 = 344
	/**/
	Num345 = 345
	/**/
	Num346 = 346
	/**/
	Num347 = 347
	/**/
	Num348 = 348
	/**/
	Num349 = 349
	/**/
	Num350 = 350
	/**/
	Num351 = 351
	/**/
	Num352 = 352
	/**/
	Num353 = 353
	/**/
	Num354 = 354
	/*By clicking "Log in" or "Sign Up," you agree to our Terms of Service and Privacy Policy*/
	Num355 = 355
	/*攻略已下架*/
	Num356 = 356
	/**/
	Num357 = 357
	/*暂无相关内容*/
	Num358 = 358
	/*生成中*/
	Num359 = 359
	/*长按保存*/
	Num360 = 360
	/*请输入您的签名*/
	Num361 = 361
	/*编辑我的个性签名*/
	Num362 = 362
	/*已上阵*/
	Num363 = 363
	/*请输入正确的区服ID*/
	Num364 = 364
	/*不可选*/
	Num365 = 365

.
)
func init() {
	for _, i18 := range I18List {
		I18Map[i18.Num] = i18
	}
}
func GetEN(lang constant.Lang, num int) string {
	i18, ok := I18Map[num]
	if !ok {
		return ""
	}
	switch lang {
	case constant.EN:
		return i18.En
	case constant.DE:
		return i18.DE
	case constant.RU:
		return i18.Ru
	case constant.ZH:
		return i18.ZH
	case constant.FR:
		return i18.FR
	default:
		return i18.ZH
	}
}

type I18 struct {
	Num int    `json:"num,omitempty" xlsx:"序号"  index:"1" `
	Img string `json:"img,omitempty" xlsx:"图片"  index:"2" `
	ZH  string `json:"zh,omitempty" xlsx:"文本"  index:"3" `
	En  string `json:"en,omitempty" xlsx:"EN"  index:"4" `
	Ru  string `json:"ru,omitempty" xlsx:"RU"  index:"5" `
	DE  string `json:"de,omitempty" xlsx:"DE"  index:"6" `
	FR  string `json:"fr,omitempty" xlsx:"FR"  index:"7" `
}

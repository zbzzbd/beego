/**
 * @preserve jQuery DateTimePicker plugin v2.4.5
 * @homepage http://xdsoft.net/jqplugins/datetimepicker/
 * (c) 2014, Chupurnov Valeriy.
 */
/*global document,window,jQuery,setTimeout,clearTimeout,HighlightedDate,getCurrentValue*/
;(function (factory) {
	if ( typeof define === 'function' && define.amd ) {
		// AMD. Register as an anonymous module.
		define(['jquery', 'jquery-mousewheel', 'date-functions'], factory);
	} else if (typeof exports === 'object') {
		// Node/CommonJS style for Browserify
		module.exports = factory;
	} else {
		// Browser globals
		factory(jQuery);
	}
}(function ($) {
	'use strict';
	var default_options  = {
		i18n: {
			ar: { // Arabic
				months: [
					"كانون الثاني", "شباط", "آذار", "نيسان", "مايو", "حزيران", "تموز", "آب", "أيلول", "تشرين الأول", "تشرين الثاني", "كانون الأول"
				],
				dayOfWeekShort: [
					"ن", "ث", "ع", "خ", "ج", "س", "ح"
				],
				dayOfWeek: ["الأحد", "الاثنين", "الثلاثاء", "الأربعاء", "الخميس", "الجمعة", "السبت", "الأحد"]
			},
			ro: { // Romanian
				months: [
					"Ianuarie", "Februarie", "Martie", "Aprilie", "Mai", "Iunie", "Iulie", "August", "Septembrie", "Octombrie", "Noiembrie", "Decembrie"
				],
				dayOfWeekShort: [
					"Du", "Lu", "Ma", "Mi", "Jo", "Vi", "Sâ"
				],
				dayOfWeek: ["Duminică", "Luni", "Marţi", "Miercuri", "Joi", "Vineri", "Sâmbătă"]
			},
			id: { // Indonesian
				months: [
					"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"
				],
				dayOfWeekShort: [
					"Min", "Sen", "Sel", "Rab", "Kam", "Jum", "Sab"
				],
				dayOfWeek: ["Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"]
			},
			is: { // Icelandic
				months: [
					"Janúar", "Febrúar", "Mars", "Apríl", "Maí", "Júní", "Júlí", "Ágúst", "September", "Október", "Nóvember", "Desember"
				],
				dayOfWeekShort: [
					"Sun", "Mán", "Þrið", "Mið", "Fim", "Fös", "Lau"
				],
				dayOfWeek: ["Sunnudagur", "Mánudagur", "Þriðjudagur", "Miðvikudagur", "Fimmtudagur", "Föstudagur", "Laugardagur"]
			},
			bg: { // Bulgarian
				months: [
					"Януари", "Февруари", "Март", "Април", "Май", "Юни", "Юли", "Август", "Септември", "Октомври", "Ноември", "Декември"
				],
				dayOfWeekShort: [
					"Нд", "Пн", "Вт", "Ср", "Чт", "Пт", "Сб"
				],
				dayOfWeek: ["Неделя", "Понеделник", "Вторник", "Сряда", "Четвъртък", "Петък", "Събота"]
			},
			fa: { // Persian/Farsi
				months: [
					'فروردین', 'اردیبهشت', 'خرداد', 'تیر', 'مرداد', 'شهریور', 'مهر', 'آبان', 'آذر', 'دی', 'بهمن', 'اسفند'
				],
				dayOfWeekShort: [
					'یکشنبه', 'دوشنبه', 'سه شنبه', 'چهارشنبه', 'پنجشنبه', 'جمعه', 'شنبه'
				],
				dayOfWeek: ["یک‌شنبه", "دوشنبه", "سه‌شنبه", "چهارشنبه", "پنج‌شنبه", "جمعه", "شنبه", "یک‌شنبه"]
			},
			ru: { // Russian
				months: [
					'Январь', 'Февраль', 'Март', 'Апрель', 'Май', 'Июнь', 'Июль', 'Август', 'Сентябрь', 'Октябрь', 'Ноябрь', 'Декабрь'
				],
				dayOfWeekShort: [
					"Вс", "Пн", "Вт", "Ср", "Чт", "Пт", "Сб"
				],
				dayOfWeek: ["Воскресенье", "Понедельник", "Вторник", "Среда", "Четверг", "Пятница", "Суббота"]
			},
			uk: { // Ukrainian
				months: [
					'Січень', 'Лютий', 'Березень', 'Квітень', 'Травень', 'Червень', 'Липень', 'Серпень', 'Вересень', 'Жовтень', 'Листопад', 'Грудень'
				],
				dayOfWeekShort: [
					"Ндл", "Пнд", "Втр", "Срд", "Чтв", "Птн", "Сбт"
				],
				dayOfWeek: ["Неділя", "Понеділок", "Вівторок", "Середа", "Четвер", "П'ятниця", "Субота"]
			},
			en: { // English
				months: [
					"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"
				],
				dayOfWeekShort: [
					"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"
				],
				dayOfWeek: ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"]
			},
			el: { // Ελληνικά
				months: [
					"Ιανουάριος", "Φεβρουάριος", "Μάρτιος", "Απρίλιος", "Μάιος", "Ιούνιος", "Ιούλιος", "Αύγουστος", "Σεπτέμβριος", "Οκτώβριος", "Νοέμβριος", "Δεκέμβριος"
				],
				dayOfWeekShort: [
					"Κυρ", "Δευ", "Τρι", "Τετ", "Πεμ", "Παρ", "Σαβ"
				],
				dayOfWeek: ["Κυριακή", "Δευτέρα", "Τρίτη", "Τετάρτη", "Πέμπτη", "Παρασκευή", "Σάββατο"]
			},
			de: { // German
				months: [
					'Januar', 'Februar', 'März', 'April', 'Mai', 'Juni', 'Juli', 'August', 'September', 'Oktober', 'November', 'Dezember'
				],
				dayOfWeekShort: [
					"So", "Mo", "Di", "Mi", "Do", "Fr", "Sa"
				],
				dayOfWeek: ["Sonntag", "Montag", "Dienstag", "Mittwoch", "Donnerstag", "Freitag", "Samstag"]
			},
			nl: { // Dutch
				months: [
					"januari", "februari", "maart", "april", "mei", "juni", "juli", "augustus", "september", "oktober", "november", "december"
				],
				dayOfWeekShort: [
					"zo", "ma", "di", "wo", "do", "vr", "za"
				],
				dayOfWeek: ["zondag", "maandag", "dinsdag", "woensdag", "donderdag", "vrijdag", "zaterdag"]
			},
			tr: { // Turkish
				months: [
					"Ocak", "Şubat", "Mart", "Nisan", "Mayıs", "Haziran", "Temmuz", "Ağustos", "Eylül", "Ekim", "Kasım", "Aralık"
				],
				dayOfWeekShort: [
					"Paz", "Pts", "Sal", "Çar", "Per", "Cum", "Cts"
				],
				dayOfWeek: ["Pazar", "Pazartesi", "Salı", "Çarşamba", "Perşembe", "Cuma", "Cumartesi"]
			},
			fr: { //French
				months: [
					"Janvier", "Février", "Mars", "Avril", "Mai", "Juin", "Juillet", "Août", "Septembre", "Octobre", "Novembre", "Décembre"
				],
				dayOfWeekShort: [
					"Dim", "Lun", "Mar", "Mer", "Jeu", "Ven", "Sam"
				],
				dayOfWeek: ["dimanche", "lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi"]
			},
			es: { // Spanish
				months: [
					"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"
				],
				dayOfWeekShort: [
					"Dom", "Lun", "Mar", "Mié", "Jue", "Vie", "Sáb"
				],
				dayOfWeek: ["Domingo", "Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado"]
			},
			th: { // Thai
				months: [
					'มกราคม', 'กุมภาพันธ์', 'มีนาคม', 'เมษายน', 'พฤษภาคม', 'มิถุนายน', 'กรกฎาคม', 'สิงหาคม', 'กันยายน', 'ตุลาคม', 'พฤศจิกายน', 'ธันวาคม'
				],
				dayOfWeekShort: [
					'อา.', 'จ.', 'อ.', 'พ.', 'พฤ.', 'ศ.', 'ส.'
				],
				dayOfWeek: ["อาทิตย์", "จันทร์", "อังคาร", "พุธ", "พฤหัส", "ศุกร์", "เสาร์", "อาทิตย์"]
			},
			pl: { // Polish
				months: [
					"styczeń", "luty", "marzec", "kwiecień", "maj", "czerwiec", "lipiec", "sierpień", "wrzesień", "październik", "listopad", "grudzień"
				],
				dayOfWeekShort: [
					"nd", "pn", "wt", "śr", "cz", "pt", "sb"
				],
				dayOfWeek: ["niedziela", "poniedziałek", "wtorek", "środa", "czwartek", "piątek", "sobota"]
			},
			pt: { // Portuguese
				months: [
					"Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho", "Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro"
				],
				dayOfWeekShort: [
					"Dom", "Seg", "Ter", "Qua", "Qui", "Sex", "Sab"
				],
				dayOfWeek: ["Domingo", "Segunda", "Terça", "Quarta", "Quinta", "Sexta", "Sábado"]
			},
			ch: { // Simplified Chinese
				months: [
					"一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"
				],
				dayOfWeekShort: [
					"日", "一", "二", "三", "四", "五", "六"
				]
			},
			se: { // Swedish
				months: [
					"Januari", "Februari", "Mars", "April", "Maj", "Juni", "Juli", "Augusti", "September",  "Oktober", "November", "December"
				],
				dayOfWeekShort: [
					"Sön", "Mån", "Tis", "Ons", "Tor", "Fre", "Lör"
				]
			},
			kr: { // Korean
				months: [
					"1월", "2월", "3월", "4월", "5월", "6월", "7월", "8월", "9월", "10월", "11월", "12월"
				],
				dayOfWeekShort: [
					"일", "월", "화", "수", "목", "금", "토"
				],
				dayOfWeek: ["일요일", "월요일", "화요일", "수요일", "목요일", "금요일", "토요일"]
			},
			it: { // Italian
				months: [
					"Gennaio", "Febbraio", "Marzo", "Aprile", "Maggio", "Giugno", "Luglio", "Agosto", "Settembre", "Ottobre", "Novembre", "Dicembre"
				],
				dayOfWeekShort: [
					"Dom", "Lun", "Mar", "Mer", "Gio", "Ven", "Sab"
				],
				dayOfWeek: ["Domenica", "Lunedì", "Martedì", "Mercoledì", "Giovedì", "Venerdì", "Sabato"]
			},
			da: { // Dansk
				months: [
					"January", "Februar", "Marts", "April", "Maj", "Juni", "July", "August", "September", "Oktober", "November", "December"
				],
				dayOfWeekShort: [
					"Søn", "Man", "Tir", "Ons", "Tor", "Fre", "Lør"
				],
				dayOfWeek: ["søndag", "mandag", "tirsdag", "onsdag", "torsdag", "fredag", "lørdag"]
			},
			no: { // Norwegian
				months: [
					"Januar", "Februar", "Mars", "April", "Mai", "Juni", "Juli", "August", "September", "Oktober", "November", "Desember"
				],
				dayOfWeekShort: [
					"Søn", "Man", "Tir", "Ons", "Tor", "Fre", "Lør"
				],
				dayOfWeek: ['Søndag', 'Mandag', 'Tirsdag', 'Onsdag', 'Torsdag', 'Fredag', 'Lørdag']
			},
			ja: { // Japanese
				months: [
					"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"
				],
				dayOfWeekShort: [
					"日", "月", "火", "水", "木", "金", "土"
				],
				dayOfWeek: ["日曜", "月曜", "火曜", "水曜", "木曜", "金曜", "土曜"]
			},
			vi: { // Vietnamese
				months: [
					"Tháng 1", "Tháng 2", "Tháng 3", "Tháng 4", "Tháng 5", "Tháng 6", "Tháng 7", "Tháng 8", "Tháng 9", "Tháng 10", "Tháng 11", "Tháng 12"
				],
				dayOfWeekShort: [
					"CN", "T2", "T3", "T4", "T5", "T6", "T7"
				],
				dayOfWeek: ["Chủ nhật", "Thứ hai", "Thứ ba", "Thứ tư", "Thứ năm", "Thứ sáu", "Thứ bảy"]
			},
			sl: { // Slovenščina
				months: [
					"Januar", "Februar", "Marec", "April", "Maj", "Junij", "Julij", "Avgust", "September", "Oktober", "November", "December"
				],
				dayOfWeekShort: [
					"Ned", "Pon", "Tor", "Sre", "Čet", "Pet", "Sob"
				],
				dayOfWeek: ["Nedelja", "Ponedeljek", "Torek", "Sreda", "Četrtek", "Petek", "Sobota"]
			},
			cs: { // Čeština
				months: [
					"Leden", "Únor", "Březen", "Duben", "Květen", "Červen", "Červenec", "Srpen", "Září", "Říjen", "Listopad", "Prosinec"
				],
				dayOfWeekShort: [
					"Ne", "Po", "Út", "St", "Čt", "Pá", "So"
				]
			},
			hu: { // Hungarian
				months: [
					"Január", "Február", "Március", "Április", "Május", "Június", "Július", "Augusztus", "Szeptember", "Október", "November", "December"
				],
				dayOfWeekShort: [
					"Va", "Hé", "Ke", "Sze", "Cs", "Pé", "Szo"
				],
				dayOfWeek: ["vasárnap", "hétfő", "kedd", "szerda", "csütörtök", "péntek", "szombat"]
			},
			az: { //Azerbaijanian (Azeri)
				months: [
					"Yanvar", "Fevral", "Mart", "Aprel", "May", "Iyun", "Iyul", "Avqust", "Sentyabr", "Oktyabr", "Noyabr", "Dekabr"
				],
				dayOfWeekShort: [
					"B", "Be", "Ça", "Ç", "Ca", "C", "Ş"
				],
				dayOfWeek: ["Bazar", "Bazar ertəsi", "Çərşənbə axşamı", "Çərşənbə", "Cümə axşamı", "Cümə", "Şənbə"]
			},
			bs: { //Bosanski
				months: [
					"Januar", "Februar", "Mart", "April", "Maj", "Jun", "Jul", "Avgust", "Septembar", "Oktobar", "Novembar", "Decembar"
				],
				dayOfWeekShort: [
					"Ned", "Pon", "Uto", "Sri", "Čet", "Pet", "Sub"
				],
				dayOfWeek: ["Nedjelja","Ponedjeljak", "Utorak", "Srijeda", "Četvrtak", "Petak", "Subota"]
			},
			ca: { //Català
				months: [
					"Gener", "Febrer", "Març", "Abril", "Maig", "Juny", "Juliol", "Agost", "Setembre", "Octubre", "Novembre", "Desembre"
				],
				dayOfWeekShort: [
					"Dg", "Dl", "Dt", "Dc", "Dj", "Dv", "Ds"
				],
				dayOfWeek: ["Diumenge", "Dilluns", "Dimarts", "Dimecres", "Dijous", "Divendres", "Dissabte"]
			},
			'en-GB': { //English (British)
				months: [
					"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"
				],
				dayOfWeekShort: [
					"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"
				],
				dayOfWeek: ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"]
			},
			et: { //"Eesti"
				months: [
					"Jaanuar", "Veebruar", "Märts", "Aprill", "Mai", "Juuni", "Juuli", "August", "September", "Oktoober", "November", "Detsember"
				],
				dayOfWeekShort: [
					"P", "E", "T", "K", "N", "R", "L"
				],
				dayOfWeek: ["Pühapäev", "Esmaspäev", "Teisipäev", "Kolmapäev", "Neljapäev", "Reede", "Laupäev"]
			},
			eu: { //Euskara
				months: [
					"Urtarrila", "Otsaila", "Martxoa", "Apirila", "Maiatza", "Ekaina", "Uztaila", "Abuztua", "Iraila", "Urria", "Azaroa", "Abendua"
				],
				dayOfWeekShort: [
					"Ig.", "Al.", "Ar.", "Az.", "Og.", "Or.", "La."
				],
				dayOfWeek: ['Igandea', 'Astelehena', 'Asteartea', 'Asteazkena', 'Osteguna', 'Ostirala', 'Larunbata']
			},
			fi: { //Finnish (Suomi)
				months: [
					"Tammikuu", "Helmikuu", "Maaliskuu", "Huhtikuu", "Toukokuu", "Kesäkuu", "Heinäkuu", "Elokuu", "Syyskuu", "Lokakuu", "Marraskuu", "Joulukuu"
				],
				dayOfWeekShort: [
					"Su", "Ma", "Ti", "Ke", "To", "Pe", "La"
				],
				dayOfWeek: ["sunnuntai", "maanantai", "tiistai", "keskiviikko", "torstai", "perjantai", "lauantai"]
			},
			gl: { //Galego
				months: [
					"Xan", "Feb", "Maz", "Abr", "Mai", "Xun", "Xul", "Ago", "Set", "Out", "Nov", "Dec"
				],
				dayOfWeekShort: [
					"Dom", "Lun", "Mar", "Mer", "Xov", "Ven", "Sab"
				],
				dayOfWeek: ["Domingo", "Luns", "Martes", "Mércores", "Xoves", "Venres", "Sábado"]
			},
			hr: { //Hrvatski
				months: [
					"Siječanj", "Veljača", "Ožujak", "Travanj", "Svibanj", "Lipanj", "Srpanj", "Kolovoz", "Rujan", "Listopad", "Studeni", "Prosinac"
				],
				dayOfWeekShort: [
					"Ned", "Pon", "Uto", "Sri", "Čet", "Pet", "Sub"
				],
				dayOfWeek: ["Nedjelja", "Ponedjeljak", "Utorak", "Srijeda", "Četvrtak", "Petak", "Subota"]
			},
			ko: { //Korean (한국어)
				months: [
					"1월", "2월", "3월", "4월", "5월", "6월", "7월", "8월", "9월", "10월", "11월", "12월"
				],
				dayOfWeekShort: [
					"일", "월", "화", "수", "목", "금", "토"
				],
				dayOfWeek: ["일요일", "월요일", "화요일", "수요일", "목요일", "금요일", "토요일"]
			},
			lt: { //Lithuanian (lietuvių)
				months: [
					"Sausio", "Vasario", "Kovo", "Balandžio", "Gegužės", "Birželio", "Liepos", "Rugpjūčio", "Rugsėjo", "Spalio", "Lapkričio", "Gruodžio"
				],
				dayOfWeekShort: [
					"Sek", "Pir", "Ant", "Tre", "Ket", "Pen", "Šeš"
				],
				dayOfWeek: ["Sekmadienis", "Pirmadienis", "Antradienis", "Trečiadienis", "Ketvirtadienis", "Penktadienis", "Šeštadienis"]
			},
			lv: { //Latvian (Latviešu)
				months: [
					"Janvāris", "Februāris", "Marts", "Aprīlis ", "Maijs", "Jūnijs", "Jūlijs", "Augusts", "Septembris", "Oktobris", "Novembris", "Decembris"
				],
				dayOfWeekShort: [
					"Sv", "Pr", "Ot", "Tr", "Ct", "Pk", "St"
				],
				dayOfWeek: ["Svētdiena", "Pirmdiena", "Otrdiena", "Trešdiena", "Ceturtdiena", "Piektdiena", "Sestdiena"]
			},
			mk: { //Macedonian (Македонски)
				months: [
					"јануари", "февруари", "март", "април", "мај", "јуни", "јули", "август", "септември", "октомври", "ноември", "декември"
				],
				dayOfWeekShort: [
					"нед", "пон", "вто", "сре", "чет", "пет", "саб"
				],
				dayOfWeek: ["Недела", "Понеделник", "Вторник", "Среда", "Четврток", "Петок", "Сабота"]
			},
			mn: { //Mongolian (Монгол)
				months: [
					"1-р сар", "2-р сар", "3-р сар", "4-р сар", "5-р сар", "6-р сар", "7-р сар", "8-р сар", "9-р сар", "10-р сар", "11-р сар", "12-р сар"
				],
				dayOfWeekShort: [
					"Дав", "Мяг", "Лха", "Пүр", "Бсн", "Бям", "Ням"
				],
				dayOfWeek: ["Даваа", "Мягмар", "Лхагва", "Пүрэв", "Баасан", "Бямба", "Ням"]
			},
			'pt-BR': { //Português(Brasil)
				months: [
					"Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho", "Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro"
				],
				dayOfWeekShort: [
					"Dom", "Seg", "Ter", "Qua", "Qui", "Sex", "Sáb"
				],
				dayOfWeek: ["Domingo", "Segunda", "Terça", "Quarta", "Quinta", "Sexta", "Sábado"]
			},
			sk: { //Slovenčina
				months: [
					"Január", "Február", "Marec", "Apríl", "Máj", "Jún", "Júl", "August", "September", "Október", "November", "December"
				],
				dayOfWeekShort: [
					"Ne", "Po", "Ut", "St", "Št", "Pi", "So"
				],
				dayOfWeek: ["Nedeľa", "Pondelok", "Utorok", "Streda", "Štvrtok", "Piatok", "Sobota"]
			},
			sq: { //Albanian (Shqip)
				months: [
					"Janar", "Shkurt", "Mars", "Prill", "Maj", "Qershor", "Korrik", "Gusht", "Shtator", "Tetor", "Nëntor", "Dhjetor"
				],
				dayOfWeekShort: [
					"Die", "Hën", "Mar", "Mër", "Enj", "Pre", "Shtu"
				],
				dayOfWeek: ["E Diel", "E Hënë", "E Martē", "E Mërkurë", "E Enjte", "E Premte", "E Shtunë"]
			},
			'sr-YU': { //Serbian (Srpski)
				months: [
					"Januar", "Februar", "Mart", "April", "Maj", "Jun", "Jul", "Avgust", "Septembar", "Oktobar", "Novembar", "Decembar"
				],
				dayOfWeekShort: [
					"Ned", "Pon", "Uto", "Sre", "čet", "Pet", "Sub"
				],
				dayOfWeek: ["Nedelja","Ponedeljak", "Utorak", "Sreda", "Četvrtak", "Petak", "Subota"]
			},
			sr: { //Serbian Cyrillic (Српски)
				months: [
					"јануар", "фебруар", "март", "април", "мај", "јун", "јул", "август", "септембар", "октобар", "новембар", "децембар"
				],
				dayOfWeekShort: [
					"нед", "пон", "уто", "сре", "чет", "пет", "суб"
				],
				dayOfWeek: ["Недеља","Понедељак", "Уторак", "Среда", "Четвртак", "Петак", "Субота"]
			},
			sv: { //Svenska
				months: [
					"Januari", "Februari", "Mars", "April", "Maj", "Juni", "Juli", "Augusti", "September", "Oktober", "November", "December"
				],
				dayOfWeekShort: [
					"Sön", "Mån", "Tis", "Ons", "Tor", "Fre", "Lör"
				],
				dayOfWeek: ["Söndag", "Måndag", "Tisdag", "Onsdag", "Torsdag", "Fredag", "Lördag"]
			},
			'zh-TW': { //Traditional Chinese (繁體中文)
				months: [
					"一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"
				],
				dayOfWeekShort: [
					"日", "一", "二", "三", "四", "五", "六"
				],
				dayOfWeek: ["星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"]
			},
			zh: { //Simplified Chinese (简体中文)
				months: [
					"一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"
				],
				dayOfWeekShort: [
					"日", "一", "二", "三", "四", "五", "六"
				],
				dayOfWeek: ["星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"]
			},
			he: { //Hebrew (עברית)
				months: [
					'ינואר', 'פברואר', 'מרץ', 'אפריל', 'מאי', 'יוני', 'יולי', 'אוגוסט', 'ספטמבר', 'אוקטובר', 'נובמבר', 'דצמבר'
				],
				dayOfWeekShort: [
					'א\'', 'ב\'', 'ג\'', 'ד\'', 'ה\'', 'ו\'', 'שבת'
				],
				dayOfWeek: ["ראשון", "שני", "שלישי", "רביעי", "חמישי", "שישי", "שבת", "ראשון"]
			},
			hy: { // Armenian
				months: [
					"Հունվար", "Փետրվար", "Մարտ", "Ապրիլ", "Մայիս", "Հունիս", "Հուլիս", "Օգոստոս", "Սեպտեմբեր", "Հոկտեմբեր", "Նոյեմբեր", "Դեկտեմբեր"
				],
				dayOfWeekShort: [
					"Կի", "Երկ", "Երք", "Չոր", "Հնգ", "Ուրբ", "Շբթ"
				],
				dayOfWeek: ["Կիրակի", "Երկուշաբթի", "Երեքշաբթի", "Չորեքշաբթի", "Հինգշաբթի", "Ուրբաթ", "Շաբաթ"]
			},
			kg: { // Kyrgyz
				months: [
					'Үчтүн айы', 'Бирдин айы', 'Жалган Куран', 'Чын Куран', 'Бугу', 'Кулжа', 'Теке', 'Баш Оона', 'Аяк Оона', 'Тогуздун айы', 'Жетинин айы', 'Бештин айы'
				],
				dayOfWeekShort: [
					"Жек", "Дүй", "Шей", "Шар", "Бей", "Жум", "Ише"
				],
				dayOfWeek: [
					"Жекшемб", "Дүйшөмб", "Шейшемб", "Шаршемб", "Бейшемби", "Жума", "Ишенб"
				]
			}
		},
		value: '',
		rtl: false,

		format:	'Y/m/d H:i',
		formatTime:	'H:i',
		formatDate:	'Y/m/d',

		startDate:	false, // new Date(), '1986/12/08', '-1970/01/05','-1970/01/05',
		step: 60,
		monthChangeSpinner: true,

		closeOnDateSelect: false,
		closeOnTimeSelect: true,
		closeOnWithoutClick: true,
		closeOnInputClick: true,

		timepicker: true,
		datepicker: true,
		weeks: false,

		defaultTime: false,	// use formatTime format (ex. '10:00' for formatTime:	'H:i')
		defaultDate: false,	// use formatDate format (ex new Date() or '1986/12/08' or '-1970/01/05' or '-1970/01/05')

		minDate: false,
		maxDate: false,
		minTime: false,
		maxTime: false,
		disabledMinTime: false,
		disabledMaxTime: false,

		allowTimes: [],
		opened: false,
		initTime: true,
		inline: false,
		theme: '',

		onSelectDate: function () {},
		onSelectTime: function () {},
		onChangeMonth: function () {},
		onChangeYear: function () {},
		onChangeDateTime: function () {},
		onShow: function () {},
		onClose: function () {},
		onGenerate: function () {},

		withoutCopyright: true,
		inverseButton: false,
		hours12: false,
		next: 'xdsoft_next',
		prev : 'xdsoft_prev',
		dayOfWeekStart: 0,
		parentID: 'body',
		timeHeightInTimePicker: 25,
		timepickerScrollbar: true,
		todayButton: true,
		prevButton: true,
		nextButton: true,
		defaultSelect: true,

		scrollMonth: true,
		scrollTime: true,
		scrollInput: true,

		lazyInit: false,
		mask: false,
		validateOnBlur: true,
		allowBlank: true,
		yearStart: 1950,
		yearEnd: 2050,
		monthStart: 0,
		monthEnd: 11,
		style: '',
		id: '',
		fixed: false,
		roundTime: 'round', // ceil, floor
		className: '',
		weekends: [],
		highlightedDates: [],
		highlightedPeriods: [],
		disabledDates : [],
		disabledWeekDays: [],
		yearOffset: 0,
		beforeShowDay: null,

		enterLikeTab: true,
		showApplyButton: false
	};

	var globalLocaleDefault = 'en',
		globalLocale = 'en';
	// for locale settings
	$.datetimepicker = {
		setLocale: function(locale){
			globalLocale = default_options.i18n[locale]?locale:globalLocaleDefault;
			// Override Parse and Format Library entities
			Date.monthNames = default_options.i18n[globalLocale].months;
			Date.dayNames = default_options.i18n[globalLocale].dayOfWeek;
		}
	};

	// fix for ie8
	if (!window.getComputedStyle) {
		window.getComputedStyle = function (el, pseudo) {
			this.el = el;
			this.getPropertyValue = function (prop) {
				var re = /(\-([a-z]){1})/g;
				if (prop === 'float') {
					prop = 'styleFloat';
				}
				if (re.test(prop)) {
					prop = prop.replace(re, function (a, b, c) {
						return c.toUpperCase();
					});
				}
				return el.currentStyle[prop] || null;
			};
			return this;
		};
	}
	if (!Array.prototype.indexOf) {
		Array.prototype.indexOf = function (obj, start) {
			var i, j;
			for (i = (start || 0), j = this.length; i < j; i += 1) {
				if (this[i] === obj) { return i; }
			}
			return -1;
		};
	}
	Date.prototype.countDaysInMonth = function () {
		return new Date(this.getFullYear(), this.getMonth() + 1, 0).getDate();
	};
	$.fn.xdsoftScroller = function (percent) {
		return this.each(function () {
			var timeboxparent = $(this),
				pointerEventToXY = function (e) {
					var out = {x: 0, y: 0},
						touch;
					if (e.type === 'touchstart' || e.type === 'touchmove' || e.type === 'touchend' || e.type === 'touchcancel') {
						touch  = e.originalEvent.touches[0] || e.originalEvent.changedTouches[0];
						out.x = touch.clientX;
						out.y = touch.clientY;
					} else if (e.type === 'mousedown' || e.type === 'mouseup' || e.type === 'mousemove' || e.type === 'mouseover' || e.type === 'mouseout' || e.type === 'mouseenter' || e.type === 'mouseleave') {
						out.x = e.clientX;
						out.y = e.clientY;
					}
					return out;
				},
				move = 0,
				timebox,
				parentHeight,
				height,
				scrollbar,
				scroller,
				maximumOffset = 100,
				start = false,
				startY = 0,
				startTop = 0,
				h1 = 0,
				touchStart = false,
				startTopScroll = 0,
				calcOffset = function () {};
			if (percent === 'hide') {
				timeboxparent.find('.xdsoft_scrollbar').hide();
				return;
			}
			if (!$(this).hasClass('xdsoft_scroller_box')) {
				timebox = timeboxparent.children().eq(0);
				parentHeight = timeboxparent[0].clientHeight;
				height = timebox[0].offsetHeight;
				scrollbar = $('<div class="xdsoft_scrollbar"></div>');
				scroller = $('<div class="xdsoft_scroller"></div>');
				scrollbar.append(scroller);

				timeboxparent.addClass('xdsoft_scroller_box').append(scrollbar);
				calcOffset = function calcOffset(event) {
					var offset = pointerEventToXY(event).y - startY + startTopScroll;
					if (offset < 0) {
						offset = 0;
					}
					if (offset + scroller[0].offsetHeight > h1) {
						offset = h1 - scroller[0].offsetHeight;
					}
					timeboxparent.trigger('scroll_element.xdsoft_scroller', [maximumOffset ? offset / maximumOffset : 0]);
				};

				scroller
					.on('touchstart.xdsoft_scroller mousedown.xdsoft_scroller', function (event) {
						if (!parentHeight) {
							timeboxparent.trigger('resize_scroll.xdsoft_scroller', [percent]);
						}

						startY = pointerEventToXY(event).y;
						startTopScroll = parseInt(scroller.css('margin-top'), 10);
						h1 = scrollbar[0].offsetHeight;

						if (event.type === 'mousedown') {
							if (document) {
								$(document.body).addClass('xdsoft_noselect');
							}
							$([document.body, window]).on('mouseup.xdsoft_scroller', function arguments_callee() {
								$([document.body, window]).off('mouseup.xdsoft_scroller', arguments_callee)
									.off('mousemove.xdsoft_scroller', calcOffset)
									.removeClass('xdsoft_noselect');
							});
							$(document.body).on('mousemove.xdsoft_scroller', calcOffset);
						} else {
							touchStart = true;
							event.stopPropagation();
							event.preventDefault();
						}
					})
					.on('touchmove', function (event) {
						if (touchStart) {
							event.preventDefault();
							calcOffset(event);
						}
					})
					.on('touchend touchcancel', function (event) {
						touchStart =  false;
						startTopScroll = 0;
					});

				timeboxparent
					.on('scroll_element.xdsoft_scroller', function (event, percentage) {
						if (!parentHeight) {
							timeboxparent.trigger('resize_scroll.xdsoft_scroller', [percentage, true]);
						}
						percentage = percentage > 1 ? 1 : (percentage < 0 || isNaN(percentage)) ? 0 : percentage;

						scroller.css('margin-top', maximumOffset * percentage);

						setTimeout(function () {
							timebox.css('marginTop', -parseInt((timebox[0].offsetHeight - parentHeight) * percentage, 10));
						}, 10);
					})
					.on('resize_scroll.xdsoft_scroller', function (event, percentage, noTriggerScroll) {
						var percent, sh;
						parentHeight = timeboxparent[0].clientHeight;
						height = timebox[0].offsetHeight;
						percent = parentHeight / height;
						sh = percent * scrollbar[0].offsetHeight;
						if (percent > 1) {
							scroller.hide();
						} else {
							scroller.show();
							scroller.css('height', parseInt(sh > 10 ? sh : 10, 10));
							maximumOffset = scrollbar[0].offsetHeight - scroller[0].offsetHeight;
							if (noTriggerScroll !== true) {
								timeboxparent.trigger('scroll_element.xdsoft_scroller', [percentage || Math.abs(parseInt(timebox.css('marginTop'), 10)) / (height - parentHeight)]);
							}
						}
					});

				timeboxparent.on('mousewheel', function (event) {
					var top = Math.abs(parseInt(timebox.css('marginTop'), 10));

					top = top - (event.deltaY * 20);
					if (top < 0) {
						top = 0;
					}

					timeboxparent.trigger('scroll_element.xdsoft_scroller', [top / (height - parentHeight)]);
					event.stopPropagation();
					return false;
				});

				timeboxparent.on('touchstart', function (event) {
					start = pointerEventToXY(event);
					startTop = Math.abs(parseInt(timebox.css('marginTop'), 10));
				});

				timeboxparent.on('touchmove', function (event) {
					if (start) {
						event.preventDefault();
						var coord = pointerEventToXY(event);
						timeboxparent.trigger('scroll_element.xdsoft_scroller', [(startTop - (coord.y - start.y)) / (height - parentHeight)]);
					}
				});

				timeboxparent.on('touchend touchcancel', function (event) {
					start = false;
					startTop = 0;
				});
			}
			timeboxparent.trigger('resize_scroll.xdsoft_scroller', [percent]);
		});
	};

	$.fn.datetimepicker = function (opt) {
		var KEY0 = 48,
			KEY9 = 57,
			_KEY0 = 96,
			_KEY9 = 105,
			CTRLKEY = 17,
			DEL = 46,
			ENTER = 13,
			ESC = 27,
			BACKSPACE = 8,
			ARROWLEFT = 37,
			ARROWUP = 38,
			ARROWRIGHT = 39,
			ARROWDOWN = 40,
			TAB = 9,
			F5 = 116,
			AKEY = 65,
			CKEY = 67,
			VKEY = 86,
			ZKEY = 90,
			YKEY = 89,
			ctrlDown	=	false,
			options = ($.isPlainObject(opt) || !opt) ? $.extend(true, {}, default_options, opt) : $.extend(true, {}, default_options),

			lazyInitTimer = 0,
			createDateTimePicker,
			destroyDateTimePicker,

			lazyInit = function (input) {
				input
					.on('open.xdsoft focusin.xdsoft mousedown.xdsoft', function initOnActionCallback(event) {
						if (input.is(':disabled') || input.data('xdsoft_datetimepicker')) {
							return;
						}
						clearTimeout(lazyInitTimer);
						lazyInitTimer = setTimeout(function () {

							if (!input.data('xdsoft_datetimepicker')) {
								createDateTimePicker(input);
							}
							input
								.off('open.xdsoft focusin.xdsoft mousedown.xdsoft', initOnActionCallback)
								.trigger('open.xdsoft');
						}, 100);
					});
			};

		createDateTimePicker = function (input) {
			var datetimepicker = $('<div class="xdsoft_datetimepicker xdsoft_noselect"></div>'),
				xdsoft_copyright = $('<div class="xdsoft_copyright"><a target="_blank" href="http://xdsoft.net/jqplugins/datetimepicker/">xdsoft.net</a></div>'),
				datepicker = $('<div class="xdsoft_datepicker active"></div>'),
				mounth_picker = $('<div class="xdsoft_mounthpicker"><button type="button" class="xdsoft_prev"></button><button type="button" class="xdsoft_today_button"></button>' +
					'<div class="xdsoft_label xdsoft_month"><span></span><i></i></div>' +
					'<div class="xdsoft_label xdsoft_year"><span></span><i></i></div>' +
					'<button type="button" class="xdsoft_next"></button></div>'),
				calendar = $('<div class="xdsoft_calendar"></div>'),
				timepicker = $('<div class="xdsoft_timepicker active"><button type="button" class="xdsoft_prev"></button><div class="xdsoft_time_box"></div><button type="button" class="xdsoft_next"></button></div>'),
				timeboxparent = timepicker.find('.xdsoft_time_box').eq(0),
				timebox = $('<div class="xdsoft_time_variant"></div>'),
				applyButton = $('<button type="button" class="xdsoft_save_selected blue-gradient-button">Save Selected</button>'),
			/*scrollbar = $('<div class="xdsoft_scrollbar"></div>'),
			 scroller = $('<div class="xdsoft_scroller"></div>'),*/
				monthselect = $('<div class="xdsoft_select xdsoft_monthselect"><div></div></div>'),
				yearselect = $('<div class="xdsoft_select xdsoft_yearselect"><div></div></div>'),
				triggerAfterOpen = false,
				XDSoft_datetime,
			//scroll_element,
				xchangeTimer,
				timerclick,
				current_time_index,
				setPos,
				timer = 0,
				timer1 = 0,
				_xdsoft_datetime;

			if (options.id) {
				datetimepicker.attr('id', options.id);
			}
			if (options.style) {
				datetimepicker.attr('style', options.style);
			}
			if (options.weeks) {
				datetimepicker.addClass('xdsoft_showweeks');
			}
			if (options.rtl) {
				datetimepicker.addClass('xdsoft_rtl');
			}

			datetimepicker.addClass('xdsoft_' + options.theme);
			datetimepicker.addClass(options.className);

			mounth_picker
				.find('.xdsoft_month span')
				.after(monthselect);
			mounth_picker
				.find('.xdsoft_year span')
				.after(yearselect);

			mounth_picker
				.find('.xdsoft_month,.xdsoft_year')
				.on('mousedown.xdsoft', function (event) {
					var select = $(this).find('.xdsoft_select').eq(0),
						val = 0,
						top = 0,
						visible = select.is(':visible'),
						items,
						i;

					mounth_picker
						.find('.xdsoft_select')
						.hide();
					if (_xdsoft_datetime.currentTime) {
						val = _xdsoft_datetime.currentTime[$(this).hasClass('xdsoft_month') ? 'getMonth' : 'getFullYear']();
					}

					select[visible ? 'hide' : 'show']();
					for (items = select.find('div.xdsoft_option'), i = 0; i < items.length; i += 1) {
						if (items.eq(i).data('value') === val) {
							break;
						} else {
							top += items[0].offsetHeight;
						}
					}

					select.xdsoftScroller(top / (select.children()[0].offsetHeight - (select[0].clientHeight)));
					event.stopPropagation();
					return false;
				});

			mounth_picker
				.find('.xdsoft_select')
				.xdsoftScroller()
				.on('mousedown.xdsoft', function (event) {
					event.stopPropagation();
					event.preventDefault();
				})
				.on('mousedown.xdsoft', '.xdsoft_option', function (event) {

					if (_xdsoft_datetime.currentTime === undefined || _xdsoft_datetime.currentTime === null) {
						_xdsoft_datetime.currentTime = _xdsoft_datetime.now();
					}

					var year = _xdsoft_datetime.currentTime.getFullYear();
					if (_xdsoft_datetime && _xdsoft_datetime.currentTime) {
						_xdsoft_datetime.currentTime[$(this).parent().parent().hasClass('xdsoft_monthselect') ? 'setMonth' : 'setFullYear']($(this).data('value'));
					}

					$(this).parent().parent().hide();

					datetimepicker.trigger('xchange.xdsoft');
					if (options.onChangeMonth && $.isFunction(options.onChangeMonth)) {
						options.onChangeMonth.call(datetimepicker, _xdsoft_datetime.currentTime, datetimepicker.data('input'));
					}

					if (year !== _xdsoft_datetime.currentTime.getFullYear() && $.isFunction(options.onChangeYear)) {
						options.onChangeYear.call(datetimepicker, _xdsoft_datetime.currentTime, datetimepicker.data('input'));
					}
				});

			datetimepicker.setOptions = function (_options) {
				var highlightedDates = {},
					getCaretPos = function (input) {
						try {
							if (document.selection && document.selection.createRange) {
								var range = document.selection.createRange();
								return range.getBookmark().charCodeAt(2) - 2;
							}
							if (input.setSelectionRange) {
								return input.selectionStart;
							}
						} catch (e) {
							return 0;
						}
					},
					setCaretPos = function (node, pos) {
						node = (typeof node === "string" || node instanceof String) ? document.getElementById(node) : node;
						if (!node) {
							return false;
						}
						if (node.createTextRange) {
							var textRange = node.createTextRange();
							textRange.collapse(true);
							textRange.moveEnd('character', pos);
							textRange.moveStart('character', pos);
							textRange.select();
							return true;
						}
						if (node.setSelectionRange) {
							node.setSelectionRange(pos, pos);
							return true;
						}
						return false;
					},
					isValidValue = function (mask, value) {
						var reg = mask
							.replace(/([\[\]\/\{\}\(\)\-\.\+]{1})/g, '\\$1')
							.replace(/_/g, '{digit+}')
							.replace(/([0-9]{1})/g, '{digit$1}')
							.replace(/\{digit([0-9]{1})\}/g, '[0-$1_]{1}')
							.replace(/\{digit[\+]\}/g, '[0-9_]{1}');
						return (new RegExp(reg)).test(value);
					};
				options = $.extend(true, {}, options, _options);

				if (_options.allowTimes && $.isArray(_options.allowTimes) && _options.allowTimes.length) {
					options.allowTimes = $.extend(true, [], _options.allowTimes);
				}

				if (_options.weekends && $.isArray(_options.weekends) && _options.weekends.length) {
					options.weekends = $.extend(true, [], _options.weekends);
				}

				if (_options.highlightedDates && $.isArray(_options.highlightedDates) && _options.highlightedDates.length) {
					$.each(_options.highlightedDates, function (index, value) {
						var splitData = $.map(value.split(','), $.trim),
							exDesc,
							hDate = new HighlightedDate(Date.parseDate(splitData[0], options.formatDate), splitData[1], splitData[2]), // date, desc, style
							keyDate = hDate.date.dateFormat(options.formatDate);
						if (highlightedDates[keyDate] !== undefined) {
							exDesc = highlightedDates[keyDate].desc;
							if (exDesc && exDesc.length && hDate.desc && hDate.desc.length) {
								highlightedDates[keyDate].desc = exDesc + "\n" + hDate.desc;
							}
						} else {
							highlightedDates[keyDate] = hDate;
						}
					});

					options.highlightedDates = $.extend(true, [], highlightedDates);
				}

				if (_options.highlightedPeriods && $.isArray(_options.highlightedPeriods) && _options.highlightedPeriods.length) {
					highlightedDates = $.extend(true, [], options.highlightedDates);
					$.each(_options.highlightedPeriods, function (index, value) {
						var dateTest, // start date
							dateEnd,
							desc,
							hDate,
							keyDate,
							exDesc,
							style;
						if ($.isArray(value)) {
							dateTest = value[0];
							dateEnd = value[1];
							desc = value[2];
							style = value[3];
						}
						else {
							var splitData = $.map(value.split(','), $.trim);
							dateTest = Date.parseDate(splitData[0], options.formatDate);
							dateEnd = Date.parseDate(splitData[1], options.formatDate);
							desc = splitData[2];
							style = splitData[3];
						}

						while (dateTest <= dateEnd) {
							hDate = new HighlightedDate(dateTest, desc, style);
							keyDate = dateTest.dateFormat(options.formatDate);
							dateTest.setDate(dateTest.getDate() + 1);
							if (highlightedDates[keyDate] !== undefined) {
								exDesc = highlightedDates[keyDate].desc;
								if (exDesc && exDesc.length && hDate.desc && hDate.desc.length) {
									highlightedDates[keyDate].desc = exDesc + "\n" + hDate.desc;
								}
							} else {
								highlightedDates[keyDate] = hDate;
							}
						}
					});

					options.highlightedDates = $.extend(true, [], highlightedDates);
				}

				if (_options.disabledDates && $.isArray(_options.disabledDates) && _options.disabledDates.length) {
					options.disabledDates = $.extend(true, [], _options.disabledDates);
				}

				if (_options.disabledWeekDays && $.isArray(_options.disabledWeekDays) && _options.disabledWeekDays.length) {
					options.disabledWeekDays = $.extend(true, [], _options.disabledWeekDays);
				}

				if ((options.open || options.opened) && (!options.inline)) {
					input.trigger('open.xdsoft');
				}

				if (options.inline) {
					triggerAfterOpen = true;
					datetimepicker.addClass('xdsoft_inline');
					input.after(datetimepicker).hide();
				}

				if (options.inverseButton) {
					options.next = 'xdsoft_prev';
					options.prev = 'xdsoft_next';
				}

				if (options.datepicker) {
					datepicker.addClass('active');
				} else {
					datepicker.removeClass('active');
				}

				if (options.timepicker) {
					timepicker.addClass('active');
				} else {
					timepicker.removeClass('active');
				}

				if (options.value) {
					_xdsoft_datetime.setCurrentTime(options.value);
					if (input && input.val) {
						input.val(_xdsoft_datetime.str);
					}
				}

				if (isNaN(options.dayOfWeekStart)) {
					options.dayOfWeekStart = 0;
				} else {
					options.dayOfWeekStart = parseInt(options.dayOfWeekStart, 10) % 7;
				}

				if (!options.timepickerScrollbar) {
					timeboxparent.xdsoftScroller('hide');
				}

				if (options.minDate && /^[\+\-](.*)$/.test(options.minDate)) {
					options.minDate = _xdsoft_datetime.strToDateTime(options.minDate).dateFormat(options.formatDate);
				}

				if (options.maxDate &&  /^[\+\-](.*)$/.test(options.maxDate)) {
					options.maxDate = _xdsoft_datetime.strToDateTime(options.maxDate).dateFormat(options.formatDate);
				}

				applyButton.toggle(options.showApplyButton);

				mounth_picker
					.find('.xdsoft_today_button')
					.css('visibility', !options.todayButton ? 'hidden' : 'visible');

				mounth_picker
					.find('.' + options.prev)
					.css('visibility', !options.prevButton ? 'hidden' : 'visible');

				mounth_picker
					.find('.' + options.next)
					.css('visibility', !options.nextButton ? 'hidden' : 'visible');

				if (options.mask) {
					input.off('keydown.xdsoft');

					if (options.mask === true) {
						options.mask = options.format
							.replace(/Y/g, '9999')
							.replace(/F/g, '9999')
							.replace(/m/g, '19')
							.replace(/d/g, '39')
							.replace(/H/g, '29')
							.replace(/i/g, '59')
							.replace(/s/g, '59');
					}

					if ($.type(options.mask) === 'string') {
						if (!isValidValue(options.mask, input.val())) {
							input.val(options.mask.replace(/[0-9]/g, '_'));
						}

						input.on('keydown.xdsoft', function (event) {
							var val = this.value,
								key = event.which,
								pos,
								digit;

							if (((key >= KEY0 && key <= KEY9) || (key >= _KEY0 && key <= _KEY9)) || (key === BACKSPACE || key === DEL)) {
								pos = getCaretPos(this);
								digit = (key !== BACKSPACE && key !== DEL) ? String.fromCharCode((_KEY0 <= key && key <= _KEY9) ? key - KEY0 : key) : '_';

								if ((key === BACKSPACE || key === DEL) && pos) {
									pos -= 1;
									digit = '_';
								}

								while (/[^0-9_]/.test(options.mask.substr(pos, 1)) && pos < options.mask.length && pos > 0) {
									pos += (key === BACKSPACE || key === DEL) ? -1 : 1;
								}

								val = val.substr(0, pos) + digit + val.substr(pos + 1);
								if ($.trim(val) === '') {
									val = options.mask.replace(/[0-9]/g, '_');
								} else {
									if (pos === options.mask.length) {
										event.preventDefault();
										return false;
									}
								}

								pos += (key === BACKSPACE || key === DEL) ? 0 : 1;
								while (/[^0-9_]/.test(options.mask.substr(pos, 1)) && pos < options.mask.length && pos > 0) {
									pos += (key === BACKSPACE || key === DEL) ? -1 : 1;
								}

								if (isValidValue(options.mask, val)) {
									this.value = val;
									setCaretPos(this, pos);
								} else if ($.trim(val) === '') {
									this.value = options.mask.replace(/[0-9]/g, '_');
								} else {
									input.trigger('error_input.xdsoft');
								}
							} else {
								if (([AKEY, CKEY, VKEY, ZKEY, YKEY].indexOf(key) !== -1 && ctrlDown) || [ESC, ARROWUP, ARROWDOWN, ARROWLEFT, ARROWRIGHT, F5, CTRLKEY, TAB, ENTER].indexOf(key) !== -1) {
									return true;
								}
							}

							event.preventDefault();
							return false;
						});
					}
				}
				if (options.validateOnBlur) {
					input
						.off('blur.xdsoft')
						.on('blur.xdsoft', function () {
							if (options.allowBlank && !$.trim($(this).val()).length) {
								$(this).val(null);
								datetimepicker.data('xdsoft_datetime').empty();
							} else if (!Date.parseDate($(this).val(), options.format)) {
								var splittedHours   = +([$(this).val()[0], $(this).val()[1]].join('')),
									splittedMinutes = +([$(this).val()[2], $(this).val()[3]].join(''));

								// parse the numbers as 0312 => 03:12
								if (!options.datepicker && options.timepicker && splittedHours >= 0 && splittedHours < 24 && splittedMinutes >= 0 && splittedMinutes < 60) {
									$(this).val([splittedHours, splittedMinutes].map(function (item) {
										return item > 9 ? item : '0' + item;
									}).join(':'));
								} else {
									$(this).val((_xdsoft_datetime.now()).dateFormat(options.format));
								}

								datetimepicker.data('xdsoft_datetime').setCurrentTime($(this).val());
							} else {
								datetimepicker.data('xdsoft_datetime').setCurrentTime($(this).val());
							}

							datetimepicker.trigger('changedatetime.xdsoft');
						});
				}
				options.dayOfWeekStartPrev = (options.dayOfWeekStart === 0) ? 6 : options.dayOfWeekStart - 1;

				datetimepicker
					.trigger('xchange.xdsoft')
					.trigger('afterOpen.xdsoft');
			};

			datetimepicker
				.data('options', options)
				.on('mousedown.xdsoft', function (event) {
					event.stopPropagation();
					event.preventDefault();
					yearselect.hide();
					monthselect.hide();
					return false;
				});

			//scroll_element = timepicker.find('.xdsoft_time_box');
			timeboxparent.append(timebox);
			timeboxparent.xdsoftScroller();

			datetimepicker.on('afterOpen.xdsoft', function () {
				timeboxparent.xdsoftScroller();
			});

			datetimepicker
				.append(datepicker)
				.append(timepicker);

			if (options.withoutCopyright !== true) {
				datetimepicker
					.append(xdsoft_copyright);
			}

			datepicker
				.append(mounth_picker)
				.append(calendar)
				.append(applyButton);

			$(options.parentID)
				.append(datetimepicker);

			XDSoft_datetime = function () {
				var _this = this;
				_this.now = function (norecursion) {
					var d = new Date(),
						date,
						time;

					if (!norecursion && options.defaultDate) {
						date = _this.strToDateTime(options.defaultDate);
						d.setFullYear(date.getFullYear());
						d.setMonth(date.getMonth());
						d.setDate(date.getDate());
					}

					if (options.yearOffset) {
						d.setFullYear(d.getFullYear() + options.yearOffset);
					}

					if (!norecursion && options.defaultTime) {
						time = _this.strtotime(options.defaultTime);
						d.setHours(time.getHours());
						d.setMinutes(time.getMinutes());
					}
					return d;
				};

				_this.isValidDate = function (d) {
					if (Object.prototype.toString.call(d) !== "[object Date]") {
						return false;
					}
					return !isNaN(d.getTime());
				};

				_this.setCurrentTime = function (dTime) {
					_this.currentTime = (typeof dTime === 'string') ? _this.strToDateTime(dTime) : _this.isValidDate(dTime) ? dTime : _this.now();
					datetimepicker.trigger('xchange.xdsoft');
				};

				_this.empty = function () {
					_this.currentTime = null;
				};

				_this.getCurrentTime = function (dTime) {
					return _this.currentTime;
				};

				_this.nextMonth = function () {

					if (_this.currentTime === undefined || _this.currentTime === null) {
						_this.currentTime = _this.now();
					}

					var month = _this.currentTime.getMonth() + 1,
						year;
					if (month === 12) {
						_this.currentTime.setFullYear(_this.currentTime.getFullYear() + 1);
						month = 0;
					}

					year = _this.currentTime.getFullYear();

					_this.currentTime.setDate(
						Math.min(
							new Date(_this.currentTime.getFullYear(), month + 1, 0).getDate(),
							_this.currentTime.getDate()
						)
					);
					_this.currentTime.setMonth(month);

					if (options.onChangeMonth && $.isFunction(options.onChangeMonth)) {
						options.onChangeMonth.call(datetimepicker, _xdsoft_datetime.currentTime, datetimepicker.data('input'));
					}

					if (year !== _this.currentTime.getFullYear() && $.isFunction(options.onChangeYear)) {
						options.onChangeYear.call(datetimepicker, _xdsoft_datetime.currentTime, datetimepicker.data('input'));
					}

					datetimepicker.trigger('xchange.xdsoft');
					return month;
				};

				_this.prevMonth = function () {

					if (_this.currentTime === undefined || _this.currentTime === null) {
						_this.currentTime = _this.now();
					}

					var month = _this.currentTime.getMonth() - 1;
					if (month === -1) {
						_this.currentTime.setFullYear(_this.currentTime.getFullYear() - 1);
						month = 11;
					}
					_this.currentTime.setDate(
						Math.min(
							new Date(_this.currentTime.getFullYear(), month + 1, 0).getDate(),
							_this.currentTime.getDate()
						)
					);
					_this.currentTime.setMonth(month);
					if (options.onChangeMonth && $.isFunction(options.onChangeMonth)) {
						options.onChangeMonth.call(datetimepicker, _xdsoft_datetime.currentTime, datetimepicker.data('input'));
					}
					datetimepicker.trigger('xchange.xdsoft');
					return month;
				};

				_this.getWeekOfYear = function (datetime) {
					var onejan = new Date(datetime.getFullYear(), 0, 1);
					return Math.ceil((((datetime - onejan) / 86400000) + onejan.getDay() + 1) / 7);
				};

				_this.strToDateTime = function (sDateTime) {
					var tmpDate = [], timeOffset, currentTime;

					if (sDateTime && sDateTime instanceof Date && _this.isValidDate(sDateTime)) {
						return sDateTime;
					}

					tmpDate = /^(\+|\-)(.*)$/.exec(sDateTime);
					if (tmpDate) {
						tmpDate[2] = Date.parseDate(tmpDate[2], options.formatDate);
					}
					if (tmpDate  && tmpDate[2]) {
						timeOffset = tmpDate[2].getTime() - (tmpDate[2].getTimezoneOffset()) * 60000;
						currentTime = new Date((_this.now(true)).getTime() + parseInt(tmpDate[1] + '1', 10) * timeOffset);
					} else {
						currentTime = sDateTime ? Date.parseDate(sDateTime, options.format) : _this.now();
					}

					if (!_this.isValidDate(currentTime)) {
						currentTime = _this.now();
					}

					return currentTime;
				};

				_this.strToDate = function (sDate) {
					if (sDate && sDate instanceof Date && _this.isValidDate(sDate)) {
						return sDate;
					}

					var currentTime = sDate ? Date.parseDate(sDate, options.formatDate) : _this.now(true);
					if (!_this.isValidDate(currentTime)) {
						currentTime = _this.now(true);
					}
					return currentTime;
				};

				_this.strtotime = function (sTime) {
					if (sTime && sTime instanceof Date && _this.isValidDate(sTime)) {
						return sTime;
					}
					var currentTime = sTime ? Date.parseDate(sTime, options.formatTime) : _this.now(true);
					if (!_this.isValidDate(currentTime)) {
						currentTime = _this.now(true);
					}
					return currentTime;
				};

				_this.str = function () {
					return _this.currentTime.dateFormat(options.format);
				};
				_this.currentTime = this.now();
			};

			_xdsoft_datetime = new XDSoft_datetime();

			applyButton.on('click', function (e) {//pathbrite
				e.preventDefault();
				datetimepicker.data('changed', true);
				_xdsoft_datetime.setCurrentTime(getCurrentValue());
				input.val(_xdsoft_datetime.str());
				datetimepicker.trigger('close.xdsoft');
			});
			mounth_picker
				.find('.xdsoft_today_button')
				.on('mousedown.xdsoft', function () {
					datetimepicker.data('changed', true);
					_xdsoft_datetime.setCurrentTime(0);
					datetimepicker.trigger('afterOpen.xdsoft');
				}).on('dblclick.xdsoft', function () {
					var currentDate = _xdsoft_datetime.getCurrentTime(), minDate, maxDate;
					currentDate = new Date(currentDate.getFullYear(), currentDate.getMonth(), currentDate.getDate());
					minDate = _xdsoft_datetime.strToDate(options.minDate);
					minDate = new Date(minDate.getFullYear(), minDate.getMonth(), minDate.getDate());
					if (currentDate < minDate) {
						return;
					}
					maxDate = _xdsoft_datetime.strToDate(options.maxDate);
					maxDate = new Date(maxDate.getFullYear(), maxDate.getMonth(), maxDate.getDate());
					if (currentDate > maxDate) {
						return;
					}
					input.val(_xdsoft_datetime.str());
					input.trigger('change');
					datetimepicker.trigger('close.xdsoft');
				});
			mounth_picker
				.find('.xdsoft_prev,.xdsoft_next')
				.on('mousedown.xdsoft', function () {
					var $this = $(this),
						timer = 0,
						stop = false;

					(function arguments_callee1(v) {
						if ($this.hasClass(options.next)) {
							_xdsoft_datetime.nextMonth();
						} else if ($this.hasClass(options.prev)) {
							_xdsoft_datetime.prevMonth();
						}
						if (options.monthChangeSpinner) {
							if (!stop) {
								timer = setTimeout(arguments_callee1, v || 100);
							}
						}
					}(500));

					$([document.body, window]).on('mouseup.xdsoft', function arguments_callee2() {
						clearTimeout(timer);
						stop = true;
						$([document.body, window]).off('mouseup.xdsoft', arguments_callee2);
					});
				});

			timepicker
				.find('.xdsoft_prev,.xdsoft_next')
				.on('mousedown.xdsoft', function () {
					var $this = $(this),
						timer = 0,
						stop = false,
						period = 110;
					(function arguments_callee4(v) {
						var pheight = timeboxparent[0].clientHeight,
							height = timebox[0].offsetHeight,
							top = Math.abs(parseInt(timebox.css('marginTop'), 10));
						if ($this.hasClass(options.next) && (height - pheight) - options.timeHeightInTimePicker >= top) {
							timebox.css('marginTop', '-' + (top + options.timeHeightInTimePicker) + 'px');
						} else if ($this.hasClass(options.prev) && top - options.timeHeightInTimePicker >= 0) {
							timebox.css('marginTop', '-' + (top - options.timeHeightInTimePicker) + 'px');
						}
						timeboxparent.trigger('scroll_element.xdsoft_scroller', [Math.abs(parseInt(timebox.css('marginTop'), 10) / (height - pheight))]);
						period = (period > 10) ? 10 : period - 10;
						if (!stop) {
							timer = setTimeout(arguments_callee4, v || period);
						}
					}(500));
					$([document.body, window]).on('mouseup.xdsoft', function arguments_callee5() {
						clearTimeout(timer);
						stop = true;
						$([document.body, window])
							.off('mouseup.xdsoft', arguments_callee5);
					});
				});

			xchangeTimer = 0;
			// base handler - generating a calendar and timepicker
			datetimepicker
				.on('xchange.xdsoft', function (event) {
					clearTimeout(xchangeTimer);
					xchangeTimer = setTimeout(function () {

						if (_xdsoft_datetime.currentTime === undefined || _xdsoft_datetime.currentTime === null) {
							_xdsoft_datetime.currentTime = _xdsoft_datetime.now();
						}

						var table =	'',
							start = new Date(_xdsoft_datetime.currentTime.getFullYear(), _xdsoft_datetime.currentTime.getMonth(), 1, 12, 0, 0),
							i = 0,
							j,
							today = _xdsoft_datetime.now(),
							maxDate = false,
							minDate = false,
							hDate,
							day,
							d,
							y,
							m,
							w,
							classes = [],
							customDateSettings,
							newRow = true,
							time = '',
							h = '',
							line_time,
							description;

						while (start.getDay() !== options.dayOfWeekStart) {
							start.setDate(start.getDate() - 1);
						}

						table += '<table><thead><tr>';

						if (options.weeks) {
							table += '<th></th>';
						}

						for (j = 0; j < 7; j += 1) {
							table += '<th>' + options.i18n[globalLocale].dayOfWeekShort[(j + options.dayOfWeekStart) % 7] + '</th>';
						}

						table += '</tr></thead>';
						table += '<tbody>';

						if (options.maxDate !== false) {
							maxDate = _xdsoft_datetime.strToDate(options.maxDate);
							maxDate = new Date(maxDate.getFullYear(), maxDate.getMonth(), maxDate.getDate(), 23, 59, 59, 999);
						}

						if (options.minDate !== false) {
							minDate = _xdsoft_datetime.strToDate(options.minDate);
							minDate = new Date(minDate.getFullYear(), minDate.getMonth(), minDate.getDate());
						}

						while (i < _xdsoft_datetime.currentTime.countDaysInMonth() || start.getDay() !== options.dayOfWeekStart || _xdsoft_datetime.currentTime.getMonth() === start.getMonth()) {
							classes = [];
							i += 1;

							day = start.getDay();
							d = start.getDate();
							y = start.getFullYear();
							m = start.getMonth();
							w = _xdsoft_datetime.getWeekOfYear(start);
							description = '';

							classes.push('xdsoft_date');

							if (options.beforeShowDay && $.isFunction(options.beforeShowDay.call)) {
								customDateSettings = options.beforeShowDay.call(datetimepicker, start);
							} else {
								customDateSettings = null;
							}

							if ((maxDate !== false && start > maxDate) || (minDate !== false && start < minDate) || (customDateSettings && customDateSettings[0] === false)) {
								classes.push('xdsoft_disabled');
							} else if (options.disabledDates.indexOf(start.dateFormat(options.formatDate)) !== -1) {
								classes.push('xdsoft_disabled');
							} else if (options.disabledWeekDays.indexOf(day) !== -1) {
								classes.push('xdsoft_disabled');
							}

							if (customDateSettings && customDateSettings[1] !== "") {
								classes.push(customDateSettings[1]);
							}

							if (_xdsoft_datetime.currentTime.getMonth() !== m) {
								classes.push('xdsoft_other_month');
							}

							if ((options.defaultSelect || datetimepicker.data('changed')) && _xdsoft_datetime.currentTime.dateFormat(options.formatDate) === start.dateFormat(options.formatDate)) {
								classes.push('xdsoft_current');
							}

							if (today.dateFormat(options.formatDate) === start.dateFormat(options.formatDate)) {
								classes.push('xdsoft_today');
							}

							if (start.getDay() === 0 || start.getDay() === 6 || options.weekends.indexOf(start.dateFormat(options.formatDate)) !== -1) {
								classes.push('xdsoft_weekend');
							}

							if (options.highlightedDates[start.dateFormat(options.formatDate)] !== undefined) {
								hDate = options.highlightedDates[start.dateFormat(options.formatDate)];
								classes.push(hDate.style === undefined ? 'xdsoft_highlighted_default' : hDate.style);
								description = hDate.desc === undefined ? '' : hDate.desc;
							}

							if (options.beforeShowDay && $.isFunction(options.beforeShowDay)) {
								classes.push(options.beforeShowDay(start));
							}

							if (newRow) {
								table += '<tr>';
								newRow = false;
								if (options.weeks) {
									table += '<th>' + w + '</th>';
								}
							}

							table += '<td data-date="' + d + '" data-month="' + m + '" data-year="' + y + '"' + ' class="xdsoft_date xdsoft_day_of_week' + start.getDay() + ' ' + classes.join(' ') + '" title="' + description + '">' +
								'<div>' + d + '</div>' +
								'</td>';

							if (start.getDay() === options.dayOfWeekStartPrev) {
								table += '</tr>';
								newRow = true;
							}

							start.setDate(d + 1);
						}
						table += '</tbody></table>';

						calendar.html(table);

						mounth_picker.find('.xdsoft_label span').eq(0).text(options.i18n[globalLocale].months[_xdsoft_datetime.currentTime.getMonth()]);
						mounth_picker.find('.xdsoft_label span').eq(1).text(_xdsoft_datetime.currentTime.getFullYear());

						// generate timebox
						time = '';
						h = '';
						m = '';

						line_time = function line_time(h, m) {
							var now = _xdsoft_datetime.now(), optionDateTime, current_time,
								isALlowTimesInit = options.allowTimes && $.isArray(options.allowTimes) && options.allowTimes.length;
							now.setHours(h);
							h = parseInt(now.getHours(), 10);
							now.setMinutes(m);
							m = parseInt(now.getMinutes(), 10);
							optionDateTime = new Date(_xdsoft_datetime.currentTime);
							optionDateTime.setHours(h);
							optionDateTime.setMinutes(m);
							classes = [];
							if ((options.minDateTime !== false && options.minDateTime > optionDateTime) || (options.maxTime !== false && _xdsoft_datetime.strtotime(options.maxTime).getTime() < now.getTime()) || (options.minTime !== false && _xdsoft_datetime.strtotime(options.minTime).getTime() > now.getTime())) {
								classes.push('xdsoft_disabled');
							}
							if ((options.minDateTime !== false && options.minDateTime > optionDateTime) || ((options.disabledMinTime !== false && now.getTime() > _xdsoft_datetime.strtotime(options.disabledMinTime).getTime()) && (options.disabledMaxTime !== false && now.getTime() < _xdsoft_datetime.strtotime(options.disabledMaxTime).getTime()))) {
								classes.push('xdsoft_disabled');
							}

							current_time = new Date(_xdsoft_datetime.currentTime);
							current_time.setHours(parseInt(_xdsoft_datetime.currentTime.getHours(), 10));

							if (!isALlowTimesInit) {
								current_time.setMinutes(Math[options.roundTime](_xdsoft_datetime.currentTime.getMinutes() / options.step) * options.step);
							}

							if ((options.initTime || options.defaultSelect || datetimepicker.data('changed')) && current_time.getHours() === parseInt(h, 10) && ((!isALlowTimesInit && options.step > 59) || current_time.getMinutes() === parseInt(m, 10))) {
								if (options.defaultSelect || datetimepicker.data('changed')) {
									classes.push('xdsoft_current');
								} else if (options.initTime) {
									classes.push('xdsoft_init_time');
								}
							}
							if (parseInt(today.getHours(), 10) === parseInt(h, 10) && parseInt(today.getMinutes(), 10) === parseInt(m, 10)) {
								classes.push('xdsoft_today');
							}
							time += '<div class="xdsoft_time ' + classes.join(' ') + '" data-hour="' + h + '" data-minute="' + m + '">' + now.dateFormat(options.formatTime) + '</div>';
						};

						if (!options.allowTimes || !$.isArray(options.allowTimes) || !options.allowTimes.length) {
							for (i = 0, j = 0; i < (options.hours12 ? 12 : 24); i += 1) {
								for (j = 0; j < 60; j += options.step) {
									h = (i < 10 ? '0' : '') + i;
									m = (j < 10 ? '0' : '') + j;
									line_time(h, m);
								}
							}
						} else {
							for (i = 0; i < options.allowTimes.length; i += 1) {
								h = _xdsoft_datetime.strtotime(options.allowTimes[i]).getHours();
								m = _xdsoft_datetime.strtotime(options.allowTimes[i]).getMinutes();
								line_time(h, m);
							}
						}

						timebox.html(time);

						opt = '';
						i = 0;

						for (i = parseInt(options.yearStart, 10) + options.yearOffset; i <= parseInt(options.yearEnd, 10) + options.yearOffset; i += 1) {
							opt += '<div class="xdsoft_option ' + (_xdsoft_datetime.currentTime.getFullYear() === i ? 'xdsoft_current' : '') + '" data-value="' + i + '">' + i + '</div>';
						}
						yearselect.children().eq(0)
							.html(opt);

						for (i = parseInt(options.monthStart, 10), opt = ''; i <= parseInt(options.monthEnd, 10); i += 1) {
							opt += '<div class="xdsoft_option ' + (_xdsoft_datetime.currentTime.getMonth() === i ? 'xdsoft_current' : '') + '" data-value="' + i + '">' + options.i18n[globalLocale].months[i] + '</div>';
						}
						monthselect.children().eq(0).html(opt);
						$(datetimepicker)
							.trigger('generate.xdsoft');
					}, 10);
					event.stopPropagation();
				})
				.on('afterOpen.xdsoft', function () {
					if (options.timepicker) {
						var classType, pheight, height, top;
						if (timebox.find('.xdsoft_current').length) {
							classType = '.xdsoft_current';
						} else if (timebox.find('.xdsoft_init_time').length) {
							classType = '.xdsoft_init_time';
						}
						if (classType) {
							pheight = timeboxparent[0].clientHeight;
							height = timebox[0].offsetHeight;
							top = timebox.find(classType).index() * options.timeHeightInTimePicker + 1;
							if ((height - pheight) < top) {
								top = height - pheight;
							}
							timeboxparent.trigger('scroll_element.xdsoft_scroller', [parseInt(top, 10) / (height - pheight)]);
						} else {
							timeboxparent.trigger('scroll_element.xdsoft_scroller', [0]);
						}
					}
				});

			timerclick = 0;
			calendar
				.on('click.xdsoft', 'td', function (xdevent) {
					xdevent.stopPropagation();  // Prevents closing of Pop-ups, Modals and Flyouts in Bootstrap
					timerclick += 1;
					var $this = $(this),
						currentTime = _xdsoft_datetime.currentTime;

					if (currentTime === undefined || currentTime === null) {
						_xdsoft_datetime.currentTime = _xdsoft_datetime.now();
						currentTime = _xdsoft_datetime.currentTime;
					}

					if ($this.hasClass('xdsoft_disabled')) {
						return false;
					}

					currentTime.setDate(1);
					currentTime.setFullYear($this.data('year'));
					currentTime.setMonth($this.data('month'));
					currentTime.setDate($this.data('date'));

					datetimepicker.trigger('select.xdsoft', [currentTime]);

					input.val(_xdsoft_datetime.str());

					if (options.onSelectDate &&	$.isFunction(options.onSelectDate)) {
						options.onSelectDate.call(datetimepicker, _xdsoft_datetime.currentTime, datetimepicker.data('input'), xdevent);
					}

					datetimepicker.data('changed', true);
					datetimepicker.trigger('xchange.xdsoft');
					datetimepicker.trigger('changedatetime.xdsoft');
					if ((timerclick > 1 || (options.closeOnDateSelect === true || (options.closeOnDateSelect === false && !options.timepicker))) && !options.inline) {
						datetimepicker.trigger('close.xdsoft');
					}
					setTimeout(function () {
						timerclick = 0;
					}, 200);
				});

			timebox
				.on('click.xdsoft', 'div', function (xdevent) {
					xdevent.stopPropagation();
					var $this = $(this),
						currentTime = _xdsoft_datetime.currentTime;

					if (currentTime === undefined || currentTime === null) {
						_xdsoft_datetime.currentTime = _xdsoft_datetime.now();
						currentTime = _xdsoft_datetime.currentTime;
					}

					if ($this.hasClass('xdsoft_disabled')) {
						return false;
					}
					currentTime.setHours($this.data('hour'));
					currentTime.setMinutes($this.data('minute'));
					datetimepicker.trigger('select.xdsoft', [currentTime]);

					datetimepicker.data('input').val(_xdsoft_datetime.str());

					if (options.onSelectTime && $.isFunction(options.onSelectTime)) {
						options.onSelectTime.call(datetimepicker, _xdsoft_datetime.currentTime, datetimepicker.data('input'), xdevent);
					}
					datetimepicker.data('changed', true);
					datetimepicker.trigger('xchange.xdsoft');
					datetimepicker.trigger('changedatetime.xdsoft');
					if (options.inline !== true && options.closeOnTimeSelect === true) {
						datetimepicker.trigger('close.xdsoft');
					}
				});


			datepicker
				.on('mousewheel.xdsoft', function (event) {
					if (!options.scrollMonth) {
						return true;
					}
					if (event.deltaY < 0) {
						_xdsoft_datetime.nextMonth();
					} else {
						_xdsoft_datetime.prevMonth();
					}
					return false;
				});

			input
				.on('mousewheel.xdsoft', function (event) {
					if (!options.scrollInput) {
						return true;
					}
					if (!options.datepicker && options.timepicker) {
						current_time_index = timebox.find('.xdsoft_current').length ? timebox.find('.xdsoft_current').eq(0).index() : 0;
						if (current_time_index + event.deltaY >= 0 && current_time_index + event.deltaY < timebox.children().length) {
							current_time_index += event.deltaY;
						}
						if (timebox.children().eq(current_time_index).length) {
							timebox.children().eq(current_time_index).trigger('mousedown');
						}
						return false;
					}
					if (options.datepicker && !options.timepicker) {
						datepicker.trigger(event, [event.deltaY, event.deltaX, event.deltaY]);
						if (input.val) {
							input.val(_xdsoft_datetime.str());
						}
						datetimepicker.trigger('changedatetime.xdsoft');
						return false;
					}
				});

			datetimepicker
				.on('changedatetime.xdsoft', function (event) {
					if (options.onChangeDateTime && $.isFunction(options.onChangeDateTime)) {
						var $input = datetimepicker.data('input');
						options.onChangeDateTime.call(datetimepicker, _xdsoft_datetime.currentTime, $input, event);
						delete options.value;
						$input.trigger('change');
					}
				})
				.on('generate.xdsoft', function () {
					if (options.onGenerate && $.isFunction(options.onGenerate)) {
						options.onGenerate.call(datetimepicker, _xdsoft_datetime.currentTime, datetimepicker.data('input'));
					}
					if (triggerAfterOpen) {
						datetimepicker.trigger('afterOpen.xdsoft');
						triggerAfterOpen = false;
					}
				})
				.on('click.xdsoft', function (xdevent) {
					xdevent.stopPropagation();
				});

			current_time_index = 0;

			setPos = function () {
				var offset = datetimepicker.data('input').offset(), top = offset.top + datetimepicker.data('input')[0].offsetHeight - 1, left = offset.left, position = "absolute", node;
				if (datetimepicker.data('input').parent().css('direction') == 'rtl')
					left -= (datetimepicker.outerWidth() - datetimepicker.data('input').outerWidth());
				if (options.fixed) {
					top -= $(window).scrollTop();
					left -= $(window).scrollLeft();
					position = "fixed";
				} else {
					if (top + datetimepicker[0].offsetHeight > $(window).height() + $(window).scrollTop()) {
						top = offset.top - datetimepicker[0].offsetHeight + 1;
					}
					if (top < 0) {
						top = 0;
					}
					if (left + datetimepicker[0].offsetWidth > $(window).width()) {
						left = $(window).width() - datetimepicker[0].offsetWidth;
					}
				}

				node = datetimepicker[0];
				do {
					node = node.parentNode;
					if (window.getComputedStyle(node).getPropertyValue('position') === 'relative' && $(window).width() >= node.offsetWidth) {
						left = left - (($(window).width() - node.offsetWidth) / 2);
						break;
					}
				} while (node.nodeName !== 'HTML');
				datetimepicker.css({
					left: left,
					top: top,
					position: position
				});
			};
			datetimepicker
				.on('open.xdsoft', function (event) {
					var onShow = true;
					if (options.onShow && $.isFunction(options.onShow)) {
						onShow = options.onShow.call(datetimepicker, _xdsoft_datetime.currentTime, datetimepicker.data('input'), event);
					}
					if (onShow !== false) {
						datetimepicker.show();
						setPos();
						$(window)
							.off('resize.xdsoft', setPos)
							.on('resize.xdsoft', setPos);

						if (options.closeOnWithoutClick) {
							$([document.body, window]).on('mousedown.xdsoft', function arguments_callee6() {
								datetimepicker.trigger('close.xdsoft');
								$([document.body, window]).off('mousedown.xdsoft', arguments_callee6);
							});
						}
					}
				})
				.on('close.xdsoft', function (event) {
					var onClose = true;
					mounth_picker
						.find('.xdsoft_month,.xdsoft_year')
						.find('.xdsoft_select')
						.hide();
					if (options.onClose && $.isFunction(options.onClose)) {
						onClose = options.onClose.call(datetimepicker, _xdsoft_datetime.currentTime, datetimepicker.data('input'), event);
					}
					if (onClose !== false && !options.opened && !options.inline) {
						datetimepicker.hide();
					}
					event.stopPropagation();
				})
				.on('toggle.xdsoft', function (event) {
					if (datetimepicker.is(':visible')) {
						datetimepicker.trigger('close.xdsoft');
					} else {
						datetimepicker.trigger('open.xdsoft');
					}
				})
				.data('input', input);

			timer = 0;
			timer1 = 0;

			datetimepicker.data('xdsoft_datetime', _xdsoft_datetime);
			datetimepicker.setOptions(options);

			function getCurrentValue() {
				var ct = false, time;

				if (options.startDate) {
					ct = _xdsoft_datetime.strToDate(options.startDate);
				} else {
					ct = options.value || ((input && input.val && input.val()) ? input.val() : '');
					if (ct) {
						ct = _xdsoft_datetime.strToDateTime(ct);
					} else if (options.defaultDate) {
						ct = _xdsoft_datetime.strToDateTime(options.defaultDate);
						if (options.defaultTime) {
							time = _xdsoft_datetime.strtotime(options.defaultTime);
							ct.setHours(time.getHours());
							ct.setMinutes(time.getMinutes());
						}
					}
				}

				if (ct && _xdsoft_datetime.isValidDate(ct)) {
					datetimepicker.data('changed', true);
				} else {
					ct = '';
				}

				return ct || 0;
			}

			_xdsoft_datetime.setCurrentTime(getCurrentValue());

			input
				.data('xdsoft_datetimepicker', datetimepicker)
				.on('open.xdsoft focusin.xdsoft mousedown.xdsoft', function (event) {
					if (input.is(':disabled') || (input.data('xdsoft_datetimepicker').is(':visible') && options.closeOnInputClick)) {
						return;
					}
					clearTimeout(timer);
					timer = setTimeout(function () {
						if (input.is(':disabled')) {
							return;
						}

						triggerAfterOpen = true;
						_xdsoft_datetime.setCurrentTime(getCurrentValue());

						datetimepicker.trigger('open.xdsoft');
					}, 100);
				})
				.on('keydown.xdsoft', function (event) {
					var val = this.value, elementSelector,
						key = event.which;
					if ([ENTER].indexOf(key) !== -1 && options.enterLikeTab) {
						elementSelector = $("input:visible,textarea:visible");
						datetimepicker.trigger('close.xdsoft');
						elementSelector.eq(elementSelector.index(this) + 1).focus();
						return false;
					}
					if ([TAB].indexOf(key) !== -1) {
						datetimepicker.trigger('close.xdsoft');
						return true;
					}
				});
		};
		destroyDateTimePicker = function (input) {
			var datetimepicker = input.data('xdsoft_datetimepicker');
			if (datetimepicker) {
				datetimepicker.data('xdsoft_datetime', null);
				datetimepicker.remove();
				input
					.data('xdsoft_datetimepicker', null)
					.off('.xdsoft');
				$(window).off('resize.xdsoft');
				$([window, document.body]).off('mousedown.xdsoft');
				if (input.unmousewheel) {
					input.unmousewheel();
				}
			}
		};
		$(document)
			.off('keydown.xdsoftctrl keyup.xdsoftctrl')
			.on('keydown.xdsoftctrl', function (e) {
				if (e.keyCode === CTRLKEY) {
					ctrlDown = true;
				}
			})
			.on('keyup.xdsoftctrl', function (e) {
				if (e.keyCode === CTRLKEY) {
					ctrlDown = false;
				}
			});
		return this.each(function () {
			var datetimepicker = $(this).data('xdsoft_datetimepicker'), $input;
			if (datetimepicker) {
				if ($.type(opt) === 'string') {
					switch (opt) {
						case 'show':
							$(this).select().focus();
							datetimepicker.trigger('open.xdsoft');
							break;
						case 'hide':
							datetimepicker.trigger('close.xdsoft');
							break;
						case 'toggle':
							datetimepicker.trigger('toggle.xdsoft');
							break;
						case 'destroy':
							destroyDateTimePicker($(this));
							break;
						case 'reset':
							this.value = this.defaultValue;
							if (!this.value || !datetimepicker.data('xdsoft_datetime').isValidDate(Date.parseDate(this.value, options.format))) {
								datetimepicker.data('changed', false);
							}
							datetimepicker.data('xdsoft_datetime').setCurrentTime(this.value);
							break;
						case 'validate':
							$input = datetimepicker.data('input');
							$input.trigger('blur.xdsoft');
							break;
					}
				} else {
					datetimepicker
						.setOptions(opt);
				}
				return 0;
			}
			if ($.type(opt) !== 'string') {
				if (!options.lazyInit || options.open || options.inline) {
					createDateTimePicker($(this));
				} else {
					lazyInit($(this));
				}
			}
		});
	};
	$.fn.datetimepicker.defaults = default_options;

	function HighlightedDate(date, desc, style) {
		"use strict";
		this.date = date;
		this.desc = desc;
		this.style = style;
	}

}));
/*!
 * jQuery Mousewheel 3.1.13
 *
 * Copyright jQuery Foundation and other contributors
 * Released under the MIT license
 * http://jquery.org/license
 */

(function (factory) {
	if ( typeof define === 'function' && define.amd ) {
		// AMD. Register as an anonymous module.
		define(['jquery'], factory);
	} else if (typeof exports === 'object') {
		// Node/CommonJS style for Browserify
		module.exports = factory;
	} else {
		// Browser globals
		factory(jQuery);
	}
}(function ($) {

	var toFix  = ['wheel', 'mousewheel', 'DOMMouseScroll', 'MozMousePixelScroll'],
		toBind = ( 'onwheel' in document || document.documentMode >= 9 ) ?
			['wheel'] : ['mousewheel', 'DomMouseScroll', 'MozMousePixelScroll'],
		slice  = Array.prototype.slice,
		nullLowestDeltaTimeout, lowestDelta;

	if ( $.event.fixHooks ) {
		for ( var i = toFix.length; i; ) {
			$.event.fixHooks[ toFix[--i] ] = $.event.mouseHooks;
		}
	}

	var special = $.event.special.mousewheel = {
		version: '3.1.12',

		setup: function() {
			if ( this.addEventListener ) {
				for ( var i = toBind.length; i; ) {
					this.addEventListener( toBind[--i], handler, false );
				}
			} else {
				this.onmousewheel = handler;
			}
			// Store the line height and page height for this particular element
			$.data(this, 'mousewheel-line-height', special.getLineHeight(this));
			$.data(this, 'mousewheel-page-height', special.getPageHeight(this));
		},

		teardown: function() {
			if ( this.removeEventListener ) {
				for ( var i = toBind.length; i; ) {
					this.removeEventListener( toBind[--i], handler, false );
				}
			} else {
				this.onmousewheel = null;
			}
			// Clean up the data we added to the element
			$.removeData(this, 'mousewheel-line-height');
			$.removeData(this, 'mousewheel-page-height');
		},

		getLineHeight: function(elem) {
			var $elem = $(elem),
				$parent = $elem['offsetParent' in $.fn ? 'offsetParent' : 'parent']();
			if (!$parent.length) {
				$parent = $('body');
			}
			return parseInt($parent.css('fontSize'), 10) || parseInt($elem.css('fontSize'), 10) || 16;
		},

		getPageHeight: function(elem) {
			return $(elem).height();
		},

		settings: {
			adjustOldDeltas: true, // see shouldAdjustOldDeltas() below
			normalizeOffset: true  // calls getBoundingClientRect for each event
		}
	};

	$.fn.extend({
		mousewheel: function(fn) {
			return fn ? this.bind('mousewheel', fn) : this.trigger('mousewheel');
		},

		unmousewheel: function(fn) {
			return this.unbind('mousewheel', fn);
		}
	});


	function handler(event) {
		var orgEvent   = event || window.event,
			args       = slice.call(arguments, 1),
			delta      = 0,
			deltaX     = 0,
			deltaY     = 0,
			absDelta   = 0,
			offsetX    = 0,
			offsetY    = 0;
		event = $.event.fix(orgEvent);
		event.type = 'mousewheel';

		// Old school scrollwheel delta
		if ( 'detail'      in orgEvent ) { deltaY = orgEvent.detail * -1;      }
		if ( 'wheelDelta'  in orgEvent ) { deltaY = orgEvent.wheelDelta;       }
		if ( 'wheelDeltaY' in orgEvent ) { deltaY = orgEvent.wheelDeltaY;      }
		if ( 'wheelDeltaX' in orgEvent ) { deltaX = orgEvent.wheelDeltaX * -1; }

		// Firefox < 17 horizontal scrolling related to DOMMouseScroll event
		if ( 'axis' in orgEvent && orgEvent.axis === orgEvent.HORIZONTAL_AXIS ) {
			deltaX = deltaY * -1;
			deltaY = 0;
		}

		// Set delta to be deltaY or deltaX if deltaY is 0 for backwards compatabilitiy
		delta = deltaY === 0 ? deltaX : deltaY;

		// New school wheel delta (wheel event)
		if ( 'deltaY' in orgEvent ) {
			deltaY = orgEvent.deltaY * -1;
			delta  = deltaY;
		}
		if ( 'deltaX' in orgEvent ) {
			deltaX = orgEvent.deltaX;
			if ( deltaY === 0 ) { delta  = deltaX * -1; }
		}

		// No change actually happened, no reason to go any further
		if ( deltaY === 0 && deltaX === 0 ) { return; }

		// Need to convert lines and pages to pixels if we aren't already in pixels
		// There are three delta modes:
		//   * deltaMode 0 is by pixels, nothing to do
		//   * deltaMode 1 is by lines
		//   * deltaMode 2 is by pages
		if ( orgEvent.deltaMode === 1 ) {
			var lineHeight = $.data(this, 'mousewheel-line-height');
			delta  *= lineHeight;
			deltaY *= lineHeight;
			deltaX *= lineHeight;
		} else if ( orgEvent.deltaMode === 2 ) {
			var pageHeight = $.data(this, 'mousewheel-page-height');
			delta  *= pageHeight;
			deltaY *= pageHeight;
			deltaX *= pageHeight;
		}

		// Store lowest absolute delta to normalize the delta values
		absDelta = Math.max( Math.abs(deltaY), Math.abs(deltaX) );

		if ( !lowestDelta || absDelta < lowestDelta ) {
			lowestDelta = absDelta;

			// Adjust older deltas if necessary
			if ( shouldAdjustOldDeltas(orgEvent, absDelta) ) {
				lowestDelta /= 40;
			}
		}

		// Adjust older deltas if necessary
		if ( shouldAdjustOldDeltas(orgEvent, absDelta) ) {
			// Divide all the things by 40!
			delta  /= 40;
			deltaX /= 40;
			deltaY /= 40;
		}

		// Get a whole, normalized value for the deltas
		delta  = Math[ delta  >= 1 ? 'floor' : 'ceil' ](delta  / lowestDelta);
		deltaX = Math[ deltaX >= 1 ? 'floor' : 'ceil' ](deltaX / lowestDelta);
		deltaY = Math[ deltaY >= 1 ? 'floor' : 'ceil' ](deltaY / lowestDelta);

		// Normalise offsetX and offsetY properties
		if ( special.settings.normalizeOffset && this.getBoundingClientRect ) {
			var boundingRect = this.getBoundingClientRect();
			offsetX = event.clientX - boundingRect.left;
			offsetY = event.clientY - boundingRect.top;
		}

		// Add information to the event object
		event.deltaX = deltaX;
		event.deltaY = deltaY;
		event.deltaFactor = lowestDelta;
		event.offsetX = offsetX;
		event.offsetY = offsetY;
		// Go ahead and set deltaMode to 0 since we converted to pixels
		// Although this is a little odd since we overwrite the deltaX/Y
		// properties with normalized deltas.
		event.deltaMode = 0;

		// Add event and delta to the front of the arguments
		args.unshift(event, delta, deltaX, deltaY);

		// Clearout lowestDelta after sometime to better
		// handle multiple device types that give different
		// a different lowestDelta
		// Ex: trackpad = 3 and mouse wheel = 120
		if (nullLowestDeltaTimeout) { clearTimeout(nullLowestDeltaTimeout); }
		nullLowestDeltaTimeout = setTimeout(nullLowestDelta, 200);

		return ($.event.dispatch || $.event.handle).apply(this, args);
	}

	function nullLowestDelta() {
		lowestDelta = null;
	}

	function shouldAdjustOldDeltas(orgEvent, absDelta) {
		// If this is an older event and the delta is divisable by 120,
		// then we are assuming that the browser is treating this as an
		// older mouse wheel event and that we should divide the deltas
		// by 40 to try and get a more usable deltaFactor.
		// Side note, this actually impacts the reported scroll distance
		// in older browsers and can cause scrolling to be slower than native.
		// Turn this off by setting $.event.special.mousewheel.settings.adjustOldDeltas to false.
		return special.settings.adjustOldDeltas && orgEvent.type === 'mousewheel' && absDelta % 120 === 0;
	}

}));
/*
 * Copyright (C) 2004 Baron Schwartz <baron at sequent dot org>
 * Modified by Jonathan Gotti aka malko <jgotti at jgotti dot org>
 *
 * This program is free software; you can redistribute it and/or modify it
 * under the terms of the GNU Lesser General Public License as published by the
 * Free Software Foundation, version 2.1.
 *
 * This program is distributed in the hope that it will be useful, but WITHOUT
 * ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
 * FOR A PARTICULAR PURPOSE.  See the GNU Lesser General Public License for more
 * details.
 */
/* jshint laxbreak:true*/
!(function (factory) {
	if ( typeof define === 'function' && define.amd ) {
		define([], factory);
	} else if (typeof exports === 'object') {
		module.exports = factory;
	} else {
		factory();
	}
}(function(){
	"use strict";
	var parseFunctions = {};
	var parseRegexes = [];
	var formatFunctions = {};
	var charFormatters = {
		d: function(date) { return stringLeftPad(date.getDate(), 2, '0'); },
		D: function(date) { return Date.dayNames[date.getDay()].substring(0, 3); },
		j: function(date) { return date.getDate(); },
		l: function(date) { return Date.dayNames[date.getDay()]; },
		S: function(date) { return date.getSuffix(); },
		w: function(date) { return date.getDay(); },
		z: function(date) { return date.getDayOfYear(); },
		W: function(date) { return date.getWeekOfYear(); },
		F: function(date) { return Date.monthNames[date.getMonth()]; },
		m: function(date) { return stringLeftPad(date.getMonth() + 1, 2, '0'); },
		M: function(date) { return Date.monthNames[date.getMonth()].substring(0, 3); },
		n: function(date) { return (date.getMonth() + 1); },
		t: function(date) { return date.getDaysInMonth(); },
		L: function(date) { return (date.isLeapYear() ? 1 : 0); },
		Y: function(date) { return date.getFullYear(); },
		y: function(date) { return ('' + date.getFullYear()).substring(2, 4); },
		a: function(date) { return (date.getHours() < 12 ? 'am' : 'pm'); },
		A: function(date) { return (date.getHours() < 12 ? 'AM' : 'PM'); },
		g: function(date) { return ((date.getHours() %12) ? date.getHours() % 12 : 12); },
		G: function(date) { return date.getHours(); },
		h: function(date) { return stringLeftPad((date.getHours() %12) ? date.getHours() % 12 : 12, 2, '0'); },
		H: function(date) { return stringLeftPad(date.getHours(), 2, '0'); },
		i: function(date) { return stringLeftPad(date.getMinutes(), 2, '0'); },
		s: function(date) { return stringLeftPad(date.getSeconds(), 2, '0'); },
		O: function(date) { return date.getGMTOffset(); },
		T: function(date) { return date.getTimezone(); },
		Z: function(date) { return (date.getTimezoneOffset() * -60); }
	};

	Date.prototype.dateFormat = function(format) {
		formatFunctions[format]  || createNewFormat(format);
		return formatFunctions[format](this);
	};

	function createNewFormat(format) {
		var formatters = [];
		var special = false;
		var ch = '';
		for (var i = 0; i < format.length; ++i) {
			ch = format.charAt(i);
			if (!special && ch === "\\") {
				special = true;
			} else if (special) {
				special = false;
				formatters.push(stringEscape(ch));
			} else {
				formatters.push(charFormatters[ch] || stringEscape(ch));
			}
		}
		formatFunctions[format] = getFormatter(formatters);
	}

	function getFormatter(formatters) {
		return function(date) {
			var res = [];
			for (var i=0, l=formatters.length; i < l; i++) {
				res.push(typeof formatters[i] === 'string' ? formatters[i] : formatters[i](date));
			}
			return res.join('');
		};
	}

	Date.parseDate = function(input, format) {
		parseFunctions[format] || createParser(format);
		return parseFunctions[format](input);
	};

	function getParser(format, regexNum, assigns) {
		return function(input){
			var d = new Date();
			var results = input.match(parseRegexes[regexNum]);
			if (results && results.length > 0) {
				results.y = d.getFullYear();
				results.m = d.getMonth();
				results.d = d.getDate();
				results.h = -1;
				results.i = -1;
				results.s = -1;
				for ( var i = 0, l = assigns.length; i < l; i++) {
					assigns[i](results);
				}
				if (results.y > 0 && results.m >= 0 && results.d > 0 && results.h >= 0 && results.i >= 0 && results.s >= 0){
					return new Date(results.y, results.m, results.d, results.h, results.i, results.s);
				} else if (results.y > 0 && results.m >= 0 && results.d > 0 && results.h >= 0 && results.i >= 0) {
					return new Date(results.y, results.m, results.d, results.h, results.i);
				} else if (results.y > 0 && results.m >= 0 && results.d > 0 && results.h >= 0) {
					return new Date(results.y, results.m, results.d, results.h);
				} else if (results.y > 0 && results.m >= 0 && results.d > 0) {
					return new Date(results.y, results.m, results.d);
				} else if (results.y > 0 && results.m >= 0) {
					return new Date(results.y, results.m);
				} else if (results.y > 0) {
					return new Date(results.y);
				}
			}
			return null;
		};
	}

	function createParser(format) {
		var regexNum = parseRegexes.length;
		var currentGroup = 1;
		var regex = "";
		var special = false;
		var ch = '';
		var assigns=[];
		var obj;
		for (var i = 0; i < format.length; ++i) {
			ch = format.charAt(i);
			if (!special && ch === "\\") {
				special = true;
			} else if (special) {
				special = false;
				regex += stringEscape(ch);
			} else {
				obj = formatCodeToRegex(ch, currentGroup);
				currentGroup += obj.g;
				regex += obj.s;
				if (obj.g && obj.a) {
					assigns.push(obj.a);
				}
			}
		}

		parseRegexes[regexNum] = new RegExp("^" + regex + "$");
		parseFunctions[format] = getParser(format, regexNum, assigns);
	}

	function formatCodeToRegex(character, currentGroup) {
		switch (character) {
			case "D":
				return {g:0,
					s:"(?:Sun|Mon|Tue|Wed|Thu|Fri|Sat)"};
			case "j":
			case "d":
				return {g:1,
					a: function(results){ results.d = parseInt(results[currentGroup], 10);},
					s:"(\\d{1,2})"
				};
			case "l":
				return {g:0,
					s:"(?:" + Date.dayNames.join("|") + ")"};
			case "S":
				return {g:0,
					s:"(?:st|nd|rd|th)"};
			case "w":
				return {g:0,
					s:"\\d"};
			case "z":
				return {g:0,
					s:"(?:\\d{1,3})"};
			case "W":
				return {g:0,
					s:"(?:\\d{2})"};
			case "F":
				return {g:1,
					a: function(results) { results.m = parseInt(Date.monthNumbers[results[currentGroup].substring(0, 3)], 10);},
					s:"(" + Date.monthNames.join("|") + ")"};
			case "M":
				return {g:1,
					a: function(results) { results.m = parseInt(Date.monthNumbers[results[currentGroup]], 10);},
					s:"(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec)"};
			case "n":
			case "m":
				return {g:1,
					a: function(results) { results.m = parseInt(results[currentGroup], 10) - 1;},
					s:"(\\d{1,2})"};
			case "t":
				return {g:0,
					s:"\\d{1,2}"};
			case "L":
				return {g:0,
					s:"(?:1|0)"};
			case "Y":
				return {g:1,
					a: function(results) { results.y = parseInt(results[currentGroup], 10);},
					s:"(\\d{4})"};
			case "y":
				return {g:1,
					a: function(results) {
						var ty = parseInt(results[currentGroup], 10);
						results.y = ty > Date.y2kYear ? 1900 + ty : 2000 + ty;
					},
					s:"(\\d{1,2})"};
			case "a":
				return {g:1,
					a: function(results) {
						if (results[currentGroup] === 'am') {
							if (results.h == 12) { results.h = 0; }
						} else {
							if (results.h < 12) { results.h += 12; }
						}
					},
					s:"(am|pm)"};
			case "A":
				return {g:1,
					a: function(results){
						if (results[currentGroup] === 'AM') {
							if (results.h == 12) { results.h = 0; }
						} else {
							if (results.h < 12) { results.h += 12; }
						}
					},
					s:"(AM|PM)"};
			case "g":
			case "G":
			case "h":
			case "H":
				return {g:1,
					a: function(results) {results.h = parseInt(results[currentGroup], 10);},
					s:"(\\d{1,2})"};
			case "i":
				return {g:1,
					a: function(results) {results.i = parseInt(results[currentGroup], 10);},
					s:"(\\d{2})"};
			case "s":
				return {g:1,
					a: function(results) {results.s = parseInt(results[currentGroup], 10);},
					s:"(\\d{2})"};
			case "O":
				return {g:0,
					s:"[+-]\\d{4}"};
			case "T":
				return {g:0,
					s:"[A-Z]{3}"};
			case "Z":
				return {g:0,
					s:"[+-]\\d{1,5}"};
			default:
				return {g:0,
					s:stringEscape(character)};
		}
	}

	Date.prototype.getTimezone = function() {
		return this.toString().replace(
			/^.*? ([A-Z]{3}) [0-9]{4}.*$/, "$1").replace(
			/^.*?\(([A-Z])[a-z]+ ([A-Z])[a-z]+ ([A-Z])[a-z]+\)$/, "$1$2$3");
	};

	Date.prototype.getGMTOffset = function() {
		return (this.getTimezoneOffset() > 0 ? "-" : "+")
			+ stringLeftPad(Math.floor(this.getTimezoneOffset() / 60), 2, "0")
			+ stringLeftPad(this.getTimezoneOffset() % 60, 2, "0");
	};

	Date.prototype.getDayOfYear = function() {
		var num = 0;
		Date.daysInMonth[1] = this.isLeapYear() ? 29 : 28;
		for (var i = 0; i < this.getMonth(); ++i) {
			num += Date.daysInMonth[i];
		}
		return num + this.getDate() - 1;
	};

	Date.prototype.getWeekOfYear = function() {
		// Skip to Thursday of this week
		var now = this.getDayOfYear() + (4 - this.getDay());
		// Find the first Thursday of the year
		var jan1 = new Date(this.getFullYear(), 0, 1);
		var then = (7 - jan1.getDay() + 4);
		return stringLeftPad(((now - then) / 7) + 1, 2, "0");
	};

	Date.prototype.isLeapYear = function() {
		var year = this.getFullYear();
		return !!((year & 3) === 0 && (year % 100 || (year % 400 === 0 && year)));
	};

	Date.prototype.getFirstDayOfMonth = function() {
		var day = (this.getDay() - (this.getDate() - 1)) % 7;
		return (day < 0) ? (day + 7) : day;
	};

	Date.prototype.getLastDayOfMonth = function() {
		var day = (this.getDay() + (Date.daysInMonth[this.getMonth()] - this.getDate())) % 7;
		return (day < 0) ? (day + 7) : day;
	};

	Date.prototype.getDaysInMonth = function() {
		Date.daysInMonth[1] = this.isLeapYear() ? 29 : 28;
		return Date.daysInMonth[this.getMonth()];
	};

	Date.prototype.getSuffix = function() {
		switch (this.getDate()) {
			case 1:
			case 21:
			case 31:
				return "st";
			case 2:
			case 22:
				return "nd";
			case 3:
			case 23:
				return "rd";
			default:
				return "th";
		}
	};

	function stringEscape(string) {
		return string.replace(/('|\\)/g, "\\$1");
	}

	function stringLeftPad(val, size, ch) {
		var result = "" + val;
		ch = ("" + ch) || " ";
		while (result.length < size) {
			result = ch + result;
		}
		return result;
	}

	Date.daysInMonth = [31,28,31,30,31,30,31,31,30,31,30,31];
	Date.monthNames =
		["January",
			"February",
			"March",
			"April",
			"May",
			"June",
			"July",
			"August",
			"September",
			"October",
			"November",
			"December"];
	Date.dayNames =
		["Sunday",
			"Monday",
			"Tuesday",
			"Wednesday",
			"Thursday",
			"Friday",
			"Saturday"];
	Date.y2kYear = 50;
	Date.monthNumbers = {
		Jan:0,
		Feb:1,
		Mar:2,
		Apr:3,
		May:4,
		Jun:5,
		Jul:6,
		Aug:7,
		Sep:8,
		Oct:9,
		Nov:10,
		Dec:11};
	Date.patterns = {
		ISO8601LongPattern:"Y-m-d H:i:s",
		ISO8601ShortPattern:"Y-m-d",
		ShortDatePattern: "n/j/Y",
		LongDatePattern: "l, F d, Y",
		FullDateTimePattern: "l, F d, Y g:i:s A",
		MonthDayPattern: "F d",
		ShortTimePattern: "g:i A",
		LongTimePattern: "g:i:s A",
		SortableDateTimePattern: "Y-m-d\\TH:i:s",
		UniversalSortableDateTimePattern: "Y-m-d H:i:sO",
		YearMonthPattern: "F, Y"};
}));
/*!
 * jQuery Validation Plugin v1.14.1-pre
 *
 * http://jqueryvalidation.org/
 *
 * Copyright (c) 2015 Jörn Zaefferer
 * Released under the MIT license
 */
(function( factory ) {
	if ( typeof define === "function" && define.amd ) {
		define( ["jquery"], factory );
	} else {
		factory( jQuery );
	}
}(function( $ ) {

	$.extend($.fn, {
		// http://jqueryvalidation.org/validate/
		validate: function( options ) {

			// if nothing is selected, return nothing; can't chain anyway
			if ( !this.length ) {
				if ( options && options.debug && window.console ) {
					console.warn( "Nothing selected, can't validate, returning nothing." );
				}
				return;
			}

			// check if a validator for this form was already created
			var validator = $.data( this[ 0 ], "validator" );
			if ( validator ) {
				return validator;
			}

			// Add novalidate tag if HTML5.
			this.attr( "novalidate", "novalidate" );

			validator = new $.validator( options, this[ 0 ] );
			$.data( this[ 0 ], "validator", validator );

			if ( validator.settings.onsubmit ) {

				this.on( "click.validate", ":submit", function( event ) {
					if ( validator.settings.submitHandler ) {
						validator.submitButton = event.target;
					}

					// allow suppressing validation by adding a cancel class to the submit button
					if ( $( this ).hasClass( "cancel" ) ) {
						validator.cancelSubmit = true;
					}

					// allow suppressing validation by adding the html5 formnovalidate attribute to the submit button
					if ( $( this ).attr( "formnovalidate" ) !== undefined ) {
						validator.cancelSubmit = true;
					}
				});

				// validate the form on submit
				this.on( "submit.validate", function( event ) {
					if ( validator.settings.debug ) {
						// prevent form submit to be able to see console output
						event.preventDefault();
					}
					function handle() {
						var hidden, result;
						if ( validator.settings.submitHandler ) {
							if ( validator.submitButton ) {
								// insert a hidden input as a replacement for the missing submit button
								hidden = $( "<input type='hidden'/>" )
									.attr( "name", validator.submitButton.name )
									.val( $( validator.submitButton ).val() )
									.appendTo( validator.currentForm );
							}
							result = validator.settings.submitHandler.call( validator, validator.currentForm, event );
							if ( validator.submitButton ) {
								// and clean up afterwards; thanks to no-block-scope, hidden can be referenced
								hidden.remove();
							}
							if ( result !== undefined ) {
								return result;
							}
							return false;
						}
						return true;
					}

					// prevent submit for invalid forms or custom submit handlers
					if ( validator.cancelSubmit ) {
						validator.cancelSubmit = false;
						return handle();
					}
					if ( validator.form() ) {
						if ( validator.pendingRequest ) {
							validator.formSubmitted = true;
							return false;
						}
						return handle();
					} else {
						validator.focusInvalid();
						return false;
					}
				});
			}

			return validator;
		},
		// http://jqueryvalidation.org/valid/
		valid: function() {
			var valid, validator, errorList;

			if ( $( this[ 0 ] ).is( "form" ) ) {
				valid = this.validate().form();
			} else {
				errorList = [];
				valid = true;
				validator = $( this[ 0 ].form ).validate();
				this.each( function() {
					valid = validator.element( this ) && valid;
					errorList = errorList.concat( validator.errorList );
				});
				validator.errorList = errorList;
			}
			return valid;
		},

		// http://jqueryvalidation.org/rules/
		rules: function( command, argument ) {
			var element = this[ 0 ],
				settings, staticRules, existingRules, data, param, filtered;

			if ( command ) {
				settings = $.data( element.form, "validator" ).settings;
				staticRules = settings.rules;
				existingRules = $.validator.staticRules( element );
				switch ( command ) {
					case "add":
						$.extend( existingRules, $.validator.normalizeRule( argument ) );
						// remove messages from rules, but allow them to be set separately
						delete existingRules.messages;
						staticRules[ element.name ] = existingRules;
						if ( argument.messages ) {
							settings.messages[ element.name ] = $.extend( settings.messages[ element.name ], argument.messages );
						}
						break;
					case "remove":
						if ( !argument ) {
							delete staticRules[ element.name ];
							return existingRules;
						}
						filtered = {};
						$.each( argument.split( /\s/ ), function( index, method ) {
							filtered[ method ] = existingRules[ method ];
							delete existingRules[ method ];
							if ( method === "required" ) {
								$( element ).removeAttr( "aria-required" );
							}
						});
						return filtered;
				}
			}

			data = $.validator.normalizeRules(
				$.extend(
					{},
					$.validator.classRules( element ),
					$.validator.attributeRules( element ),
					$.validator.dataRules( element ),
					$.validator.staticRules( element )
				), element );

			// make sure required is at front
			if ( data.required ) {
				param = data.required;
				delete data.required;
				data = $.extend( { required: param }, data );
				$( element ).attr( "aria-required", "true" );
			}

			// make sure remote is at back
			if ( data.remote ) {
				param = data.remote;
				delete data.remote;
				data = $.extend( data, { remote: param });
			}

			return data;
		}
	});

// Custom selectors
	$.extend( $.expr[ ":" ], {
		// http://jqueryvalidation.org/blank-selector/
		blank: function( a ) {
			return !$.trim( "" + $( a ).val() );
		},
		// http://jqueryvalidation.org/filled-selector/
		filled: function( a ) {
			return !!$.trim( "" + $( a ).val() );
		},
		// http://jqueryvalidation.org/unchecked-selector/
		unchecked: function( a ) {
			return !$( a ).prop( "checked" );
		}
	});

// constructor for validator
	$.validator = function( options, form ) {
		this.settings = $.extend( true, {}, $.validator.defaults, options );
		this.currentForm = form;
		this.init();
	};

// http://jqueryvalidation.org/jQuery.validator.format/
	$.validator.format = function( source, params ) {
		if ( arguments.length === 1 ) {
			return function() {
				var args = $.makeArray( arguments );
				args.unshift( source );
				return $.validator.format.apply( this, args );
			};
		}
		if ( arguments.length > 2 && params.constructor !== Array  ) {
			params = $.makeArray( arguments ).slice( 1 );
		}
		if ( params.constructor !== Array ) {
			params = [ params ];
		}
		$.each( params, function( i, n ) {
			source = source.replace( new RegExp( "\\{" + i + "\\}", "g" ), function() {
				return n;
			});
		});
		return source;
	};

	$.extend( $.validator, {

		defaults: {
			messages: {},
			groups: {},
			rules: {},
			errorClass: "error",
			validClass: "valid",
			errorElement: "label",
			focusCleanup: false,
			focusInvalid: true,
			errorContainer: $( [] ),
			errorLabelContainer: $( [] ),
			onsubmit: true,
			ignore: ":hidden",
			ignoreTitle: false,
			onfocusin: function( element ) {
				this.lastActive = element;

				// Hide error label and remove error class on focus if enabled
				if ( this.settings.focusCleanup ) {
					if ( this.settings.unhighlight ) {
						this.settings.unhighlight.call( this, element, this.settings.errorClass, this.settings.validClass );
					}
					this.hideThese( this.errorsFor( element ) );
				}
			},
			onfocusout: function( element ) {
				if ( !this.checkable( element ) && ( element.name in this.submitted || !this.optional( element ) ) ) {
					this.element( element );
				}
			},
			onkeyup: function( element, event ) {
				// Avoid revalidate the field when pressing one of the following keys
				// Shift       => 16
				// Ctrl        => 17
				// Alt         => 18
				// Caps lock   => 20
				// End         => 35
				// Home        => 36
				// Left arrow  => 37
				// Up arrow    => 38
				// Right arrow => 39
				// Down arrow  => 40
				// Insert      => 45
				// Num lock    => 144
				// AltGr key   => 225
				var excludedKeys = [
					16, 17, 18, 20, 35, 36, 37,
					38, 39, 40, 45, 144, 225
				];

				if ( event.which === 9 && this.elementValue( element ) === "" || $.inArray( event.keyCode, excludedKeys ) !== -1 ) {
					return;
				} else if ( element.name in this.submitted || element === this.lastElement ) {
					this.element( element );
				}
			},
			onclick: function( element ) {
				// click on selects, radiobuttons and checkboxes
				if ( element.name in this.submitted ) {
					this.element( element );

					// or option elements, check parent select in that case
				} else if ( element.parentNode.name in this.submitted ) {
					this.element( element.parentNode );
				}
			},
			highlight: function( element, errorClass, validClass ) {
				if ( element.type === "radio" ) {
					this.findByName( element.name ).addClass( errorClass ).removeClass( validClass );
				} else {
					$( element ).addClass( errorClass ).removeClass( validClass );
				}
			},
			unhighlight: function( element, errorClass, validClass ) {
				if ( element.type === "radio" ) {
					this.findByName( element.name ).removeClass( errorClass ).addClass( validClass );
				} else {
					$( element ).removeClass( errorClass ).addClass( validClass );
				}
			}
		},

		// http://jqueryvalidation.org/jQuery.validator.setDefaults/
		setDefaults: function( settings ) {
			$.extend( $.validator.defaults, settings );
		},

		messages: {
			required: "请输入数据。必填。",
			remote: "请修正此字段。",
			email: "请输入有效的电子邮件地址。",
			url: "请输入有效的网址。",
			date: "请输入有效的日期。",
			dateISO: "请输入有效的日期 (YYYY-MM-DD)。",
			number: "请输入有效的数字(含负数，小数) 。",
			digits: "请输入有效的整数。",
			creditcard: "请输入有效的信用卡号码。",
			equalTo: "请再次输入相同的数据.",
			maxlength: $.validator.format( "最多只能输入 {0} 个字符." ),
			minlength: $.validator.format( "请输入至少 {0} 个字符" ),
			rangelength: $.validator.format( "Please enter a value between {0} and {1} characters long." ),
			range: $.validator.format( "Please enter a value between {0} and {1}." ),
			max: $.validator.format( "此项可输入最大值为 {0}." ),
			min: $.validator.format( "此项可输入最小值为 {0}." )
		},

		autoCreateRanges: false,

		prototype: {

			init: function() {
				this.labelContainer = $( this.settings.errorLabelContainer );
				this.errorContext = this.labelContainer.length && this.labelContainer || $( this.currentForm );
				this.containers = $( this.settings.errorContainer ).add( this.settings.errorLabelContainer );
				this.submitted = {};
				this.valueCache = {};
				this.pendingRequest = 0;
				this.pending = {};
				this.invalid = {};
				this.reset();

				var groups = ( this.groups = {} ),
					rules;
				$.each( this.settings.groups, function( key, value ) {
					if ( typeof value === "string" ) {
						value = value.split( /\s/ );
					}
					$.each( value, function( index, name ) {
						groups[ name ] = key;
					});
				});
				rules = this.settings.rules;
				$.each( rules, function( key, value ) {
					rules[ key ] = $.validator.normalizeRule( value );
				});

				function delegate( event ) {
					var validator = $.data( this.form, "validator" ),
						eventType = "on" + event.type.replace( /^validate/, "" ),
						settings = validator.settings;
					if ( settings[ eventType ] && !$( this ).is( settings.ignore ) ) {
						settings[ eventType ].call( validator, this, event );
					}
				}

				$( this.currentForm )
					.on( "focusin.validate focusout.validate keyup.validate",
					":text, [type='password'], [type='file'], select, textarea, [type='number'], [type='search'], " +
					"[type='tel'], [type='url'], [type='email'], [type='datetime'], [type='date'], [type='month'], " +
					"[type='week'], [type='time'], [type='datetime-local'], [type='range'], [type='color'], " +
					"[type='radio'], [type='checkbox']", delegate)
					// Support: Chrome, oldIE
					// "select" is provided as event.target when clicking a option
					.on("click.validate", "select, option, [type='radio'], [type='checkbox']", delegate);

				if ( this.settings.invalidHandler ) {
					$( this.currentForm ).on( "invalid-form.validate", this.settings.invalidHandler );
				}

				// Add aria-required to any Static/Data/Class required fields before first validation
				// Screen readers require this attribute to be present before the initial submission http://www.w3.org/TR/WCAG-TECHS/ARIA2.html
				$( this.currentForm ).find( "[required], [data-rule-required], .required" ).attr( "aria-required", "true" );
			},

			// http://jqueryvalidation.org/Validator.form/
			form: function() {
				this.checkForm();
				$.extend( this.submitted, this.errorMap );
				this.invalid = $.extend({}, this.errorMap );
				if ( !this.valid() ) {
					$( this.currentForm ).triggerHandler( "invalid-form", [ this ]);
				}
				this.showErrors();
				return this.valid();
			},

			checkForm: function() {
				this.prepareForm();
				for ( var i = 0, elements = ( this.currentElements = this.elements() ); elements[ i ]; i++ ) {
					this.check( elements[ i ] );
				}
				return this.valid();
			},

			// http://jqueryvalidation.org/Validator.element/
			element: function( element ) {
				var cleanElement = this.clean( element ),
					checkElement = this.validationTargetFor( cleanElement ),
					result = true;

				this.lastElement = checkElement;

				if ( checkElement === undefined ) {
					delete this.invalid[ cleanElement.name ];
				} else {
					this.prepareElement( checkElement );
					this.currentElements = $( checkElement );

					result = this.check( checkElement ) !== false;
					if ( result ) {
						delete this.invalid[ checkElement.name ];
					} else {
						this.invalid[ checkElement.name ] = true;
					}
				}
				// Add aria-invalid status for screen readers
				$( element ).attr( "aria-invalid", !result );

				if ( !this.numberOfInvalids() ) {
					// Hide error containers on last error
					this.toHide = this.toHide.add( this.containers );
				}
				this.showErrors();
				return result;
			},

			// http://jqueryvalidation.org/Validator.showErrors/
			showErrors: function( errors ) {
				if ( errors ) {
					// add items to error list and map
					$.extend( this.errorMap, errors );
					this.errorList = [];
					for ( var name in errors ) {
						this.errorList.push({
							message: errors[ name ],
							element: this.findByName( name )[ 0 ]
						});
					}
					// remove items from success list
					this.successList = $.grep( this.successList, function( element ) {
						return !( element.name in errors );
					});
				}
				if ( this.settings.showErrors ) {
					this.settings.showErrors.call( this, this.errorMap, this.errorList );
				} else {
					this.defaultShowErrors();
				}
			},

			// http://jqueryvalidation.org/Validator.resetForm/
			resetForm: function() {
				if ( $.fn.resetForm ) {
					$( this.currentForm ).resetForm();
				}
				this.submitted = {};
				this.lastElement = null;
				this.prepareForm();
				this.hideErrors();
				var i, elements = this.elements()
					.removeData( "previousValue" )
					.removeAttr( "aria-invalid" );

				if ( this.settings.unhighlight ) {
					for ( i = 0; elements[ i ]; i++ ) {
						this.settings.unhighlight.call( this, elements[ i ],
							this.settings.errorClass, "" );
					}
				} else {
					elements.removeClass( this.settings.errorClass );
				}
			},

			numberOfInvalids: function() {
				return this.objectLength( this.invalid );
			},

			objectLength: function( obj ) {
				/* jshint unused: false */
				var count = 0,
					i;
				for ( i in obj ) {
					count++;
				}
				return count;
			},

			hideErrors: function() {
				this.hideThese( this.toHide );
			},

			hideThese: function( errors ) {
				errors.not( this.containers ).text( "" );
				this.addWrapper( errors ).hide();
			},

			valid: function() {
				return this.size() === 0;
			},

			size: function() {
				return this.errorList.length;
			},

			focusInvalid: function() {
				if ( this.settings.focusInvalid ) {
					try {
						$( this.findLastActive() || this.errorList.length && this.errorList[ 0 ].element || [])
							.filter( ":visible" )
							.focus()
							// manually trigger focusin event; without it, focusin handler isn't called, findLastActive won't have anything to find
							.trigger( "focusin" );
					} catch ( e ) {
						// ignore IE throwing errors when focusing hidden elements
					}
				}
			},

			findLastActive: function() {
				var lastActive = this.lastActive;
				return lastActive && $.grep( this.errorList, function( n ) {
						return n.element.name === lastActive.name;
					}).length === 1 && lastActive;
			},

			elements: function() {
				var validator = this,
					rulesCache = {};

				// select all valid inputs inside the form (no submit or reset buttons)
				return $( this.currentForm )
					.find( "input, select, textarea" )
					.not( ":submit, :reset, :image, :disabled" )
					.not( this.settings.ignore )
					.filter( function() {
						if ( !this.name && validator.settings.debug && window.console ) {
							console.error( "%o has no name assigned", this );
						}

						// select only the first element for each name, and only those with rules specified
						if ( this.name in rulesCache || !validator.objectLength( $( this ).rules() ) ) {
							return false;
						}

						rulesCache[ this.name ] = true;
						return true;
					});
			},

			clean: function( selector ) {
				return $( selector )[ 0 ];
			},

			errors: function() {
				var errorClass = this.settings.errorClass.split( " " ).join( "." );
				return $( this.settings.errorElement + "." + errorClass, this.errorContext );
			},

			reset: function() {
				this.successList = [];
				this.errorList = [];
				this.errorMap = {};
				this.toShow = $( [] );
				this.toHide = $( [] );
				this.currentElements = $( [] );
			},

			prepareForm: function() {
				this.reset();
				this.toHide = this.errors().add( this.containers );
			},

			prepareElement: function( element ) {
				this.reset();
				this.toHide = this.errorsFor( element );
			},

			elementValue: function( element ) {
				var val,
					$element = $( element ),
					type = element.type;

				if ( type === "radio" || type === "checkbox" ) {
					return this.findByName( element.name ).filter(":checked").val();
				} else if ( type === "number" && typeof element.validity !== "undefined" ) {
					return element.validity.badInput ? false : $element.val();
				}

				val = $element.val();
				if ( typeof val === "string" ) {
					return val.replace(/\r/g, "" );
				}
				return val;
			},

			check: function( element ) {
				element = this.validationTargetFor( this.clean( element ) );

				var rules = $( element ).rules(),
					rulesCount = $.map( rules, function( n, i ) {
						return i;
					}).length,
					dependencyMismatch = false,
					val = this.elementValue( element ),
					result, method, rule;

				for ( method in rules ) {
					rule = { method: method, parameters: rules[ method ] };
					try {

						result = $.validator.methods[ method ].call( this, val, element, rule.parameters );

						// if a method indicates that the field is optional and therefore valid,
						// don't mark it as valid when there are no other rules
						if ( result === "dependency-mismatch" && rulesCount === 1 ) {
							dependencyMismatch = true;
							continue;
						}
						dependencyMismatch = false;

						if ( result === "pending" ) {
							this.toHide = this.toHide.not( this.errorsFor( element ) );
							return;
						}

						if ( !result ) {
							this.formatAndAdd( element, rule );
							return false;
						}
					} catch ( e ) {
						if ( this.settings.debug && window.console ) {
							console.log( "Exception occurred when checking element " + element.id + ", check the '" + rule.method + "' method.", e );
						}
						if ( e instanceof TypeError ) {
							e.message += ".  Exception occurred when checking element " + element.id + ", check the '" + rule.method + "' method.";
						}

						throw e;
					}
				}
				if ( dependencyMismatch ) {
					return;
				}
				if ( this.objectLength( rules ) ) {
					this.successList.push( element );
				}
				return true;
			},

			// return the custom message for the given element and validation method
			// specified in the element's HTML5 data attribute
			// return the generic message if present and no method specific message is present
			customDataMessage: function( element, method ) {
				return $( element ).data( "msg" + method.charAt( 0 ).toUpperCase() +
						method.substring( 1 ).toLowerCase() ) || $( element ).data( "msg" );
			},

			// return the custom message for the given element name and validation method
			customMessage: function( name, method ) {
				var m = this.settings.messages[ name ];
				return m && ( m.constructor === String ? m : m[ method ]);
			},

			// return the first defined argument, allowing empty strings
			findDefined: function() {
				for ( var i = 0; i < arguments.length; i++) {
					if ( arguments[ i ] !== undefined ) {
						return arguments[ i ];
					}
				}
				return undefined;
			},

			defaultMessage: function( element, method ) {
				return this.findDefined(
					this.customMessage( element.name, method ),
					this.customDataMessage( element, method ),
					// title is never undefined, so handle empty string as undefined
					!this.settings.ignoreTitle && element.title || undefined,
					$.validator.messages[ method ],
					"<strong>Warning: No message defined for " + element.name + "</strong>"
				);
			},

			formatAndAdd: function( element, rule ) {
				var message = this.defaultMessage( element, rule.method ),
					theregex = /\$?\{(\d+)\}/g;
				if ( typeof message === "function" ) {
					message = message.call( this, rule.parameters, element );
				} else if ( theregex.test( message ) ) {
					message = $.validator.format( message.replace( theregex, "{$1}" ), rule.parameters );
				}
				this.errorList.push({
					message: message,
					element: element,
					method: rule.method
				});

				this.errorMap[ element.name ] = message;
				this.submitted[ element.name ] = message;
			},

			addWrapper: function( toToggle ) {
				if ( this.settings.wrapper ) {
					toToggle = toToggle.add( toToggle.parent( this.settings.wrapper ) );
				}
				return toToggle;
			},

			defaultShowErrors: function() {
				var i, elements, error;
				for ( i = 0; this.errorList[ i ]; i++ ) {
					error = this.errorList[ i ];
					if ( this.settings.highlight ) {
						this.settings.highlight.call( this, error.element, this.settings.errorClass, this.settings.validClass );
					}
					this.showLabel( error.element, error.message );
				}
				if ( this.errorList.length ) {
					this.toShow = this.toShow.add( this.containers );
				}
				if ( this.settings.success ) {
					for ( i = 0; this.successList[ i ]; i++ ) {
						this.showLabel( this.successList[ i ] );
					}
				}
				if ( this.settings.unhighlight ) {
					for ( i = 0, elements = this.validElements(); elements[ i ]; i++ ) {
						this.settings.unhighlight.call( this, elements[ i ], this.settings.errorClass, this.settings.validClass );
					}
				}
				this.toHide = this.toHide.not( this.toShow );
				this.hideErrors();
				this.addWrapper( this.toShow ).show();
			},

			validElements: function() {
				return this.currentElements.not( this.invalidElements() );
			},

			invalidElements: function() {
				return $( this.errorList ).map(function() {
					return this.element;
				});
			},

			showLabel: function( element, message ) {
				var place, group, errorID,
					error = this.errorsFor( element ),
					elementID = this.idOrName( element ),
					describedBy = $( element ).attr( "aria-describedby" );
				if ( error.length ) {
					// refresh error/success class
					error.removeClass( this.settings.validClass ).addClass( this.settings.errorClass );
					// replace message on existing label
					error.html( message );
				} else {
					// create error element
					error = $( "<" + this.settings.errorElement + ">" )
						.attr( "id", elementID + "-error" )
						.addClass( this.settings.errorClass )
						.html( message || "" );

					// Maintain reference to the element to be placed into the DOM
					place = error;
					if ( this.settings.wrapper ) {
						// make sure the element is visible, even in IE
						// actually showing the wrapped element is handled elsewhere
						place = error.hide().show().wrap( "<" + this.settings.wrapper + "/>" ).parent();
					}
					if ( this.labelContainer.length ) {
						this.labelContainer.append( place );
					} else if ( this.settings.errorPlacement ) {
						this.settings.errorPlacement( place, $( element ) );
					} else {
						place.insertAfter( element );
					}

					// Link error back to the element
					if ( error.is( "label" ) ) {
						// If the error is a label, then associate using 'for'
						error.attr( "for", elementID );
					} else if ( error.parents( "label[for='" + elementID + "']" ).length === 0 ) {
						// If the element is not a child of an associated label, then it's necessary
						// to explicitly apply aria-describedby

						errorID = error.attr( "id" ).replace( /(:|\.|\[|\]|\$)/g, "\\$1");
						// Respect existing non-error aria-describedby
						if ( !describedBy ) {
							describedBy = errorID;
						} else if ( !describedBy.match( new RegExp( "\\b" + errorID + "\\b" ) ) ) {
							// Add to end of list if not already present
							describedBy += " " + errorID;
						}
						$( element ).attr( "aria-describedby", describedBy );

						// If this element is grouped, then assign to all elements in the same group
						group = this.groups[ element.name ];
						if ( group ) {
							$.each( this.groups, function( name, testgroup ) {
								if ( testgroup === group ) {
									$( "[name='" + name + "']", this.currentForm )
										.attr( "aria-describedby", error.attr( "id" ) );
								}
							});
						}
					}
				}
				if ( !message && this.settings.success ) {
					error.text( "" );
					if ( typeof this.settings.success === "string" ) {
						error.addClass( this.settings.success );
					} else {
						this.settings.success( error, element );
					}
				}
				this.toShow = this.toShow.add( error );
			},

			errorsFor: function( element ) {
				var name = this.idOrName( element ),
					describer = $( element ).attr( "aria-describedby" ),
					selector = "label[for='" + name + "'], label[for='" + name + "'] *";

				// aria-describedby should directly reference the error element
				if ( describer ) {
					selector = selector + ", #" + describer.replace( /\s+/g, ", #" );
				}
				return this
					.errors()
					.filter( selector );
			},

			idOrName: function( element ) {
				return this.groups[ element.name ] || ( this.checkable( element ) ? element.name : element.id || element.name );
			},

			validationTargetFor: function( element ) {

				// If radio/checkbox, validate first element in group instead
				if ( this.checkable( element ) ) {
					element = this.findByName( element.name );
				}

				// Always apply ignore filter
				return $( element ).not( this.settings.ignore )[ 0 ];
			},

			checkable: function( element ) {
				return ( /radio|checkbox/i ).test( element.type );
			},

			findByName: function( name ) {
				return $( this.currentForm ).find( "[name='" + name + "']" );
			},

			getLength: function( value, element ) {
				switch ( element.nodeName.toLowerCase() ) {
					case "select":
						return $( "option:selected", element ).length;
					case "input":
						if ( this.checkable( element ) ) {
							return this.findByName( element.name ).filter( ":checked" ).length;
						}
				}
				return value.length;
			},

			depend: function( param, element ) {
				return this.dependTypes[typeof param] ? this.dependTypes[typeof param]( param, element ) : true;
			},

			dependTypes: {
				"boolean": function( param ) {
					return param;
				},
				"string": function( param, element ) {
					return !!$( param, element.form ).length;
				},
				"function": function( param, element ) {
					return param( element );
				}
			},

			optional: function( element ) {
				var val = this.elementValue( element );
				return !$.validator.methods.required.call( this, val, element ) && "dependency-mismatch";
			},

			startRequest: function( element ) {
				if ( !this.pending[ element.name ] ) {
					this.pendingRequest++;
					this.pending[ element.name ] = true;
				}
			},

			stopRequest: function( element, valid ) {
				this.pendingRequest--;
				// sometimes synchronization fails, make sure pendingRequest is never < 0
				if ( this.pendingRequest < 0 ) {
					this.pendingRequest = 0;
				}
				delete this.pending[ element.name ];
				if ( valid && this.pendingRequest === 0 && this.formSubmitted && this.form() ) {
					$( this.currentForm ).submit();
					this.formSubmitted = false;
				} else if (!valid && this.pendingRequest === 0 && this.formSubmitted ) {
					$( this.currentForm ).triggerHandler( "invalid-form", [ this ]);
					this.formSubmitted = false;
				}
			},

			previousValue: function( element ) {
				return $.data( element, "previousValue" ) || $.data( element, "previousValue", {
						old: null,
						valid: true,
						message: this.defaultMessage( element, "remote" )
					});
			},

			// cleans up all forms and elements, removes validator-specific events
			destroy: function() {
				this.resetForm();

				$( this.currentForm )
					.off( ".validate" )
					.removeData( "validator" );
			}

		},

		classRuleSettings: {
			required: { required: true },
			email: { email: true },
			url: { url: true },
			date: { date: true },
			dateISO: { dateISO: true },
			number: { number: true },
			digits: { digits: true },
			creditcard: { creditcard: true }
		},

		addClassRules: function( className, rules ) {
			if ( className.constructor === String ) {
				this.classRuleSettings[ className ] = rules;
			} else {
				$.extend( this.classRuleSettings, className );
			}
		},

		classRules: function( element ) {
			var rules = {},
				classes = $( element ).attr( "class" );

			if ( classes ) {
				$.each( classes.split( " " ), function() {
					if ( this in $.validator.classRuleSettings ) {
						$.extend( rules, $.validator.classRuleSettings[ this ]);
					}
				});
			}
			return rules;
		},

		normalizeAttributeRule: function( rules, type, method, value ) {

			// convert the value to a number for number inputs, and for text for backwards compability
			// allows type="date" and others to be compared as strings
			if ( /min|max/.test( method ) && ( type === null || /number|range|text/.test( type ) ) ) {
				value = Number( value );

				// Support Opera Mini, which returns NaN for undefined minlength
				if ( isNaN( value ) ) {
					value = undefined;
				}
			}

			if ( value || value === 0 ) {
				rules[ method ] = value;
			} else if ( type === method && type !== "range" ) {

				// exception: the jquery validate 'range' method
				// does not test for the html5 'range' type
				rules[ method ] = true;
			}
		},

		attributeRules: function( element ) {
			var rules = {},
				$element = $( element ),
				type = element.getAttribute( "type" ),
				method, value;

			for ( method in $.validator.methods ) {

				// support for <input required> in both html5 and older browsers
				if ( method === "required" ) {
					value = element.getAttribute( method );

					// Some browsers return an empty string for the required attribute
					// and non-HTML5 browsers might have required="" markup
					if ( value === "" ) {
						value = true;
					}

					// force non-HTML5 browsers to return bool
					value = !!value;
				} else {
					value = $element.attr( method );
				}

				this.normalizeAttributeRule( rules, type, method, value );
			}

			// maxlength may be returned as -1, 2147483647 ( IE ) and 524288 ( safari ) for text inputs
			if ( rules.maxlength && /-1|2147483647|524288/.test( rules.maxlength ) ) {
				delete rules.maxlength;
			}

			return rules;
		},

		dataRules: function( element ) {
			var rules = {},
				$element = $( element ),
				type = element.getAttribute( "type" ),
				method, value;

			for ( method in $.validator.methods ) {
				value = $element.data( "rule" + method.charAt( 0 ).toUpperCase() + method.substring( 1 ).toLowerCase() );
				this.normalizeAttributeRule( rules, type, method, value );
			}
			return rules;
		},

		staticRules: function( element ) {
			var rules = {},
				validator = $.data( element.form, "validator" );

			if ( validator.settings.rules ) {
				rules = $.validator.normalizeRule( validator.settings.rules[ element.name ] ) || {};
			}
			return rules;
		},

		normalizeRules: function( rules, element ) {
			// handle dependency check
			$.each( rules, function( prop, val ) {
				// ignore rule when param is explicitly false, eg. required:false
				if ( val === false ) {
					delete rules[ prop ];
					return;
				}
				if ( val.param || val.depends ) {
					var keepRule = true;
					switch ( typeof val.depends ) {
						case "string":
							keepRule = !!$( val.depends, element.form ).length;
							break;
						case "function":
							keepRule = val.depends.call( element, element );
							break;
					}
					if ( keepRule ) {
						rules[ prop ] = val.param !== undefined ? val.param : true;
					} else {
						delete rules[ prop ];
					}
				}
			});

			// evaluate parameters
			$.each( rules, function( rule, parameter ) {
				rules[ rule ] = $.isFunction( parameter ) ? parameter( element ) : parameter;
			});

			// clean number parameters
			$.each([ "minlength", "maxlength" ], function() {
				if ( rules[ this ] ) {
					rules[ this ] = Number( rules[ this ] );
				}
			});
			$.each([ "rangelength", "range" ], function() {
				var parts;
				if ( rules[ this ] ) {
					if ( $.isArray( rules[ this ] ) ) {
						rules[ this ] = [ Number( rules[ this ][ 0 ]), Number( rules[ this ][ 1 ] ) ];
					} else if ( typeof rules[ this ] === "string" ) {
						parts = rules[ this ].replace(/[\[\]]/g, "" ).split( /[\s,]+/ );
						rules[ this ] = [ Number( parts[ 0 ]), Number( parts[ 1 ] ) ];
					}
				}
			});

			if ( $.validator.autoCreateRanges ) {
				// auto-create ranges
				if ( rules.min != null && rules.max != null ) {
					rules.range = [ rules.min, rules.max ];
					delete rules.min;
					delete rules.max;
				}
				if ( rules.minlength != null && rules.maxlength != null ) {
					rules.rangelength = [ rules.minlength, rules.maxlength ];
					delete rules.minlength;
					delete rules.maxlength;
				}
			}

			return rules;
		},

		// Converts a simple string to a {string: true} rule, e.g., "required" to {required:true}
		normalizeRule: function( data ) {
			if ( typeof data === "string" ) {
				var transformed = {};
				$.each( data.split( /\s/ ), function() {
					transformed[ this ] = true;
				});
				data = transformed;
			}
			return data;
		},

		// http://jqueryvalidation.org/jQuery.validator.addMethod/
		addMethod: function( name, method, message ) {
			$.validator.methods[ name ] = method;
			$.validator.messages[ name ] = message !== undefined ? message : $.validator.messages[ name ];
			if ( method.length < 3 ) {
				$.validator.addClassRules( name, $.validator.normalizeRule( name ) );
			}
		},

		methods: {

			// http://jqueryvalidation.org/required-method/
			required: function( value, element, param ) {
				// check if dependency is met
				if ( !this.depend( param, element ) ) {
					return "dependency-mismatch";
				}
				if ( element.nodeName.toLowerCase() === "select" ) {
					// could be an array for select-multiple or a string, both are fine this way
					var val = $( element ).val();
					return val && val.length > 0;
				}
				if ( this.checkable( element ) ) {
					return this.getLength( value, element ) > 0;
				}
				return value.length > 0;
			},

			// http://jqueryvalidation.org/email-method/
			email: function( value, element ) {
				// From https://html.spec.whatwg.org/multipage/forms.html#valid-e-mail-address
				// Retrieved 2014-01-14
				// If you have a problem with this implementation, report a bug against the above spec
				// Or use custom methods to implement your own email validation
				return this.optional( element ) || /^[a-zA-Z0-9.!#$%&'*+\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$/.test( value );
			},

			// http://jqueryvalidation.org/url-method/
			url: function( value, element ) {

				// Copyright (c) 2010-2013 Diego Perini, MIT licensed
				// https://gist.github.com/dperini/729294
				// see also https://mathiasbynens.be/demo/url-regex
				// modified to allow protocol-relative URLs
				return this.optional( element ) || /^(?:(?:(?:https?|ftp):)?\/\/)(?:\S+(?::\S*)?@)?(?:(?!(?:10|127)(?:\.\d{1,3}){3})(?!(?:169\.254|192\.168)(?:\.\d{1,3}){2})(?!172\.(?:1[6-9]|2\d|3[0-1])(?:\.\d{1,3}){2})(?:[1-9]\d?|1\d\d|2[01]\d|22[0-3])(?:\.(?:1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.(?:[1-9]\d?|1\d\d|2[0-4]\d|25[0-4]))|(?:(?:[a-z\u00a1-\uffff0-9]-*)*[a-z\u00a1-\uffff0-9]+)(?:\.(?:[a-z\u00a1-\uffff0-9]-*)*[a-z\u00a1-\uffff0-9]+)*(?:\.(?:[a-z\u00a1-\uffff]{2,})).?)(?::\d{2,5})?(?:[/?#]\S*)?$/i.test( value );
			},

			// http://jqueryvalidation.org/date-method/
			date: function( value, element ) {
				return this.optional( element ) || !/Invalid|NaN/.test( new Date( value ).toString() );
			},

			// http://jqueryvalidation.org/dateISO-method/
			dateISO: function( value, element ) {
				return this.optional( element ) || /^\d{4}[\/\-](0?[1-9]|1[012])[\/\-](0?[1-9]|[12][0-9]|3[01])$/.test( value );
			},

			// http://jqueryvalidation.org/number-method/
			number: function( value, element ) {
				return this.optional( element ) || /^(?:-?\d+|-?\d{1,3}(?:,\d{3})+)?(?:\.\d+)?$/.test( value );
			},

			// http://jqueryvalidation.org/digits-method/
			digits: function( value, element ) {
				return this.optional( element ) || /^\d+$/.test( value );
			},

			// http://jqueryvalidation.org/creditcard-method/
			// based on http://en.wikipedia.org/wiki/Luhn_algorithm
			creditcard: function( value, element ) {
				if ( this.optional( element ) ) {
					return "dependency-mismatch";
				}
				// accept only spaces, digits and dashes
				if ( /[^0-9 \-]+/.test( value ) ) {
					return false;
				}
				var nCheck = 0,
					nDigit = 0,
					bEven = false,
					n, cDigit;

				value = value.replace( /\D/g, "" );

				// Basing min and max length on
				// http://developer.ean.com/general_info/Valid_Credit_Card_Types
				if ( value.length < 13 || value.length > 19 ) {
					return false;
				}

				for ( n = value.length - 1; n >= 0; n--) {
					cDigit = value.charAt( n );
					nDigit = parseInt( cDigit, 10 );
					if ( bEven ) {
						if ( ( nDigit *= 2 ) > 9 ) {
							nDigit -= 9;
						}
					}
					nCheck += nDigit;
					bEven = !bEven;
				}

				return ( nCheck % 10 ) === 0;
			},

			// http://jqueryvalidation.org/minlength-method/
			minlength: function( value, element, param ) {
				var length = $.isArray( value ) ? value.length : this.getLength( value, element );
				return this.optional( element ) || length >= param;
			},

			// http://jqueryvalidation.org/maxlength-method/
			maxlength: function( value, element, param ) {
				var length = $.isArray( value ) ? value.length : this.getLength( value, element );
				return this.optional( element ) || length <= param;
			},

			// http://jqueryvalidation.org/rangelength-method/
			rangelength: function( value, element, param ) {
				var length = $.isArray( value ) ? value.length : this.getLength( value, element );
				return this.optional( element ) || ( length >= param[ 0 ] && length <= param[ 1 ] );
			},

			// http://jqueryvalidation.org/min-method/
			min: function( value, element, param ) {
				return this.optional( element ) || value >= param;
			},

			// http://jqueryvalidation.org/max-method/
			max: function( value, element, param ) {
				return this.optional( element ) || value <= param;
			},

			// http://jqueryvalidation.org/range-method/
			range: function( value, element, param ) {
				return this.optional( element ) || ( value >= param[ 0 ] && value <= param[ 1 ] );
			},

			// http://jqueryvalidation.org/equalTo-method/
			equalTo: function( value, element, param ) {
				// bind to the blur event of the target in order to revalidate whenever the target field is updated
				// TODO find a way to bind the event just once, avoiding the unbind-rebind overhead
				var target = $( param );
				if ( this.settings.onfocusout ) {
					target.off( ".validate-equalTo" ).on( "blur.validate-equalTo", function() {
						$( element ).valid();
					});
				}
				return value === target.val();
			},

			// http://jqueryvalidation.org/remote-method/
			remote: function( value, element, param ) {
				if ( this.optional( element ) ) {
					return "dependency-mismatch";
				}

				var previous = this.previousValue( element ),
					validator, data, optionDataString;

				if (!this.settings.messages[ element.name ] ) {
					this.settings.messages[ element.name ] = {};
				}
				previous.originalMessage = this.settings.messages[ element.name ].remote;
				this.settings.messages[ element.name ].remote = previous.message;

				param = typeof param === "string" && { url: param } || param;
				optionDataString = $.param( $.extend( { data: value }, param.data ) );
				if (previous.old === optionDataString) {
					return previous.valid;
				}

				previous.old = optionDataString;
				validator = this;
				this.startRequest( element );
				data = {};
				data[ element.name ] = value;
				$.ajax( $.extend( true, {
					mode: "abort",
					port: "validate" + element.name,
					dataType: "json",
					data: data,
					context: validator.currentForm,
					success: function( response ) {
						var valid = response === true || response === "true",
							errors, message, submitted;

						validator.settings.messages[ element.name ].remote = previous.originalMessage;
						if ( valid ) {
							submitted = validator.formSubmitted;
							validator.prepareElement( element );
							validator.formSubmitted = submitted;
							validator.successList.push( element );
							delete validator.invalid[ element.name ];
							validator.showErrors();
						} else {
							errors = {};
							message = response || validator.defaultMessage( element, "remote" );
							errors[ element.name ] = previous.message = $.isFunction( message ) ? message( value ) : message;
							validator.invalid[ element.name ] = true;
							validator.showErrors( errors );
						}
						previous.valid = valid;
						validator.stopRequest( element, valid );
					}
				}, param ) );
				return "pending";
			}
		}

	});

// ajax mode: abort
// usage: $.ajax({ mode: "abort"[, port: "uniqueport"]});
// if mode:"abort" is used, the previous request on that port (port can be undefined) is aborted via XMLHttpRequest.abort()

	var pendingRequests = {},
		ajax;
// Use a prefilter if available (1.5+)
	if ( $.ajaxPrefilter ) {
		$.ajaxPrefilter(function( settings, _, xhr ) {
			var port = settings.port;
			if ( settings.mode === "abort" ) {
				if ( pendingRequests[port] ) {
					pendingRequests[port].abort();
				}
				pendingRequests[port] = xhr;
			}
		});
	} else {
		// Proxy ajax
		ajax = $.ajax;
		$.ajax = function( settings ) {
			var mode = ( "mode" in settings ? settings : $.ajaxSettings ).mode,
				port = ( "port" in settings ? settings : $.ajaxSettings ).port;
			if ( mode === "abort" ) {
				if ( pendingRequests[port] ) {
					pendingRequests[port].abort();
				}
				pendingRequests[port] = ajax.apply(this, arguments);
				return pendingRequests[port];
			}
			return ajax.apply(this, arguments);
		};
	}

}));
$(function() {
    $('#job_assign_form').form({
        fields: {
            result: {
                identifier: 'to_user',
                rules: [{
                    type: 'empty',
                    prompt: '转发给谁不能为空'
                }]
            }
        },
        onSuccess:function(){
            job_assign()
            return false;
        }
    });
});

function job_assign() {
    $.ajax({
        url:"/produce/job/assign",
        data: $("#job_assign_form").serialize(),
        type:"post",
        success:function(data){
            if (data && data.error) {
                $("#job_assign_form").removeClass("success").addClass("error")
                $("#job_assign_form #error_message").append("<li>"+data.error+"</li>");
            } else {
                window.location.href =  "/produce/job/claim";
            }
        }
    });
}

$(function(){
    function removeUsers() {
        $('#employee option').each(function(){
           if ($(this).val()) $(this).remove();
        });
    }

    var users = $('#employee option');
    function setUsers (val) {
        users.each(function(){
            if ($("#department  option:selected").attr("tag")) {
                if ($(this).attr('tag') && $(this).attr('tag').indexOf($("#department  option:selected").attr("tag")) > -1) {
                    $('#employee').append($(this));
                }
            }

            if (!val) {
                $('#employee').closest('.field').addClass('disabled');
            } else {
                $('#employee').closest('.field').removeClass('disabled');
            }
        });
    }

    $('#department').dropdown({
            placeholder: '请选择',
            onChange: function(value, text, $selectedItem) {
                removeUsers();
                setUsers(value);
            }
        })
    ;
    $('.ui.search.dropdown').dropdown({fullTextSearch:true})
    removeUsers();
    setUsers($("#department option:selected").text());

    $('#projectProgress').dropdown({placeholder:'优先级1为最高'});

    $('#projectBussiness').dropdown({placeholder:'请选择业务担当'});
    $('#projectArt').dropdown({placeholder:'请选择美术单元'});
    $('#projectTech').dropdown({placeholder:'请选择技术单元'});
    
    $('select.dropdown').dropdown({placeholder:'请选择'});
    $('.datetimepicker').datetimepicker({
        format:'Y-m-d H:i', step:5
    });
    $.datetimepicker.setLocale('zh');
    $('.datetimepicker').datetimepicker({format:'Y-m-d H:i', step:5});
}());

function selectFiles (id) {
    $('#'+id).click();
}

function setUploadFiles (e, id) {
    var strFiles = '';
    var files = $(e).prop('files');

    for(var i=0,len=files.length; i<len; i++){
        strFiles += '<label style="line-height: 24px;" class="pr20">' + files[i].name + '</label>';
    }

    $('#'+id+' .file.text').html(strFiles);

    if (files.length > 0) {
        $('#'+id).show();
    } else {
        $('#'+id).hide();
    }
}

function clearForm (formSel) {
    $(formSel).form('clear');
    $(formSel).submit();
}

function submitForm (formSel) {
    $(formSel).submit();
}


function getJob(code, callback) {
    $.ajax({
        url:"/search/job",
        data: {
            id : code
        },                               
        type:"get",
        success:function(data){
            if (data && data.error) {
               alert("查询失败：" + data.error)
            }else { 
              callback(data);

            }
        }
    });
}

function deleteComplaint(id) {
    $('.ui.modal.delete').modal({
        closable  : false,
        onDeny: function(){
        },
        onApprove: function() {
            $.ajax({
                url:"/job/complaint/del/" + id,
                type:"get",
                success:function(data){
                    if (data && data.error) {
                        alert(data.error);
                    } else {
                        window.location.reload();
                    }
                }
            });
        }
    }).modal('show');
}
function jobAssign(jobId, callback) {
    $.ajax({
        url: "/job/assign",
        data: {
            id: jobId,
            employee_id: $.cookie('uid')
        },
        type: "get",
        success: function(data) {
            if (data && data.error) {
                alert("认领错误：" + data.error)
            } else {
                window.location.reload();
            }
        }
    });
}

function delJobFile(fileId, callback) {
    $.ajax({
        url: "/job/file/del",
        data: {
            id: fileId
        },
        type: "get",
        success: function(data) {
            if (data && data.error) {
                alert("错误：" + data.error)
            } else {
                window.location.reload();
            }
        }
    });
}
var totalSize = 0;
var i =0;
function del_files(obj,id) {
    if (document.getElementById(id).files && document.getElementById(id).files[0]!=null) { 
         size= document.getElementById(id).files[0].size
        if (totalSize >0){
            totalSize -= size;  
        }
      } 
    $(obj).parent(".field").remove()
     
    return false
}

function upload_files(obj) {
    
    if (document.getElementById('fileToUpload'+i).files) {
        size= document.getElementById('fileToUpload'+i).files[0].size
        totalSize += size;  
        if (totalSize >= 1024*1024*10) {
            alert("文件已经大于10MB，请重新选择")
            totalSize -= size;
        } 
      } else {
        alert("上传文件获取为空")
      }  
}

$(function() {
    $("#add_files").click(function(){
        i=i+1; 
        s='fileToUpload'+i 
        var template = $('<div class="field"> <input type="file" name="files[]" onchange="upload_files(this)" id="'+s+'"> <button class="ui red button" onclick="del_files(this,s)">删除附件</button> </div>')  
        $(this).parent(".field").after(template) 
        
        return false
    }); 
});
   
function delJob(id) {
    $('.ui.modal.delete').modal({closable: false, onDeny: function(){}, onApprove: function(){
        $.ajax({
            url:"/job/delete/" +id,
            type: "get",
            success: function(data){
                if (data && data.error) {
                    alert(data.error);
                }else {
                    window.location.reload();
                }
            }
        });

    }}).modal('show');
}

function recoverJob(id) {
    $('.ui.modal.delete').modal({closable: false, onDeny: function(){}, onApprove: function(){
        $.ajax({
            url:"/job/recover/" +id,
            type: "get",
            success: function(data){
                if (data && data.error) {
                    alert(data.error);
                }else {
                    window.location.reload();
                }
            }
        });

    }}).modal('show');
}
$(function() {
    $('#login_form').form({
        on: 'blur',
        fields: {
            email: {
                identifier  : 'email',
                rules: [
                {
                    type   : 'empty',
                    prompt : '请输入邮箱'
                },
                {
                    type: 'email',
                    prompt: "邮箱格式不正确"
                }
                ]
            },
            password: {
                identifier  : 'password',
                rules: [
                {
                    type   : 'empty',
                    prompt : '请输入密码'
                },
                {
                    type   : 'length[6]',
                    prompt : '密码至少6位'
                }
                ]
            }
        },
        onSuccess:function(){
            login();
            return false;
        }
    });
});

function login() {
    $.ajax({
        url:"/login",
        data: $("#login_form").serialize(),
        type:"post",
        success:function(data){
            if (data && data.error) {
                $("#login_form").removeClass("success").addClass("error")
                $("#login_error_msg ul").append("<li>"+data.error+"</li>");
            } else {
                window.location.href =  "/";
            }
        }
    });
}

function initMenu () {
    var curUrl = window.location.pathname;
    if (curUrl.indexOf('?') > -1) {
        curUrl = curUrl.substr(0, curUrl.indexOf('?'))
    }
    $(".main-menu .ui.menu .item").each(function(){
        if (curUrl === '/') {
            if ($(this).attr('href') === '/') $(this).addClass("active");
        }
        else if (curUrl.indexOf($(this).attr('href')) === 0 && $(this).attr('href') !== '/') {
            $(this).addClass("active");
        }
        else if (curUrl.indexOf($(this).attr('href-tag')) === 0 ) {
            $(this).addClass("active");
        }
        if (curUrl.indexOf("/project/edit")===0 && $(this).attr('href') === '/project/list'){
            $(this).addClass("active");
        }
        if (curUrl.indexOf("/job/view")===0 && $(this).attr('href') === '/job/progress'){
            $(this).addClass("active");
        }
    });
}

(function($){
    initMenu();
})(jQuery);

;
$(function() {
    $('#project_search').search({
        source: window.projectNames,
        searchFullText: true
    });

    $('#project_create_form')
        .form({
            fields: {
                started: {
                    identifier: 'started',
                    rules: [{
                        type: 'empty',
                        prompt: '请选择启动时间'
                    }]
                },
                name: {
                    identifier: 'name',
                    rules: [{
                        type: 'empty',
                        prompt: '请填写项目名称'
                    }]
                },
                bussiness_user: {
                    identifier: 'bussiness_user',
                    rules: [{
                        type: 'empty',
                        prompt: '请选择业务担当'
                    }]
                },
                art_user: {
                    identifier: 'art_user',
                    rules: [{
                        type: 'empty',
                        prompt: '请选择美术单元'
                    }]
                },
                tech_user: {
                    identifier: 'tech_user',
                    rules: [{
                        type: 'empty',
                        prompt: '请选择技术单元'
                    }]
                },
            },
            onSuccess: function() {
                $.ajax({
                    url: "/project/create",
                    data: new FormData($("#project_create_form")[0]),
                    async: false,
                    cache: false,
                    contentType: false,
                    processData: false,
                    type: "post",
                    success: function(data) {
                        if (data && data.result != 0) {
                            $(".ui.error.message").html(data.error);
                        } else if (data.id) {
                            $(".ui.error.message").html("");
                            window.location.href = '/project/list';
                        } else {
                            $(".ui.error.message").html("未知错误");
                        }
                    }
                });
                return false;
            }
        });
});
 
function delProject(id) {
    $('.ui.modal.delete').modal({
        closable  : false,
        onDeny: function(){
        },
        onApprove: function() {
            $.ajax({
                url:"/project/del/" + id,
                type:"get",
                success:function(data){
                    if (data && data.error) {
                        alert(data.error);
                    } else {
                        window.location.reload();
                    }
                }
            });
        }
    }).modal('show');
}

function addRequire() {
    $('.ui.modal.edit').modal({
        closable  : false,
        onDeny: function(){
        },
        onApprove: function() {
            $.ajax({
                url:"/require/create",
                data: {name: $("#modal-input").val()},
                type:"get",
                success:function(data){
                    if (data && data.error) {
                        alert(data.error);
                    } else {
                        window.location.reload();
                    }
                }
            });
        }
    }).modal('show');
}

function editRequire(name) {
    $("#modal-input").val(name);

    $('.ui.modal.edit').modal({
        closable  : false,
        onDeny: function(){
        },
        onApprove: function() {
            $.ajax({
                url:"/require/edit",
                data: {name: name, newName: $("#modal-input").val()},
                type:"get",
                success:function(data){
                    if (data && data.error) {
                        alert(data.error);
                    } else {
                        window.location.reload();
                    }
                }
            });
        }
    }).modal('show');
}


function deleteRequire(name) {
    $('.ui.modal.delete').modal({
        closable  : false,
        onDeny: function(){
        },
        onApprove: function() {
            $.ajax({
                url:"/require/delete",
                data: {name: name},
                type:"get",
                success:function(data){
                    if (data && data.error) {
                        alert(data.error);
                    } else {
                        window.location.reload();
                    }
                }
            });
        }
    }).modal('show');
}

function filterRequire(obj) {
    var txtFind = $(obj).val().toLowerCase();

    $("#require-list>tr").each(function(){
        var txt = $(this).find(".name").text().toLowerCase();
        if (txt.indexOf(txtFind) < 0) {
            $(this).hide();
        } else {
            $(this).show();
        }
    });
}

$(function() {
    $('.ui.form.user').form({
        on: 'blur',
        fields: {
            company: {
                identifier  : 'company',
                rules: [
                {
                    type   : 'empty',
                    prompt : '请选择公司名称'
                }
                ]
            },
            role: {
                identifier  : 'role',
                rules: [
                    {
                        type   : 'empty',
                        prompt : '请选择角色'
                    }
                ]
            },
            name: {
                identifier  : 'name',
                rules: [
                    {
                        type   : 'empty',
                        prompt : '请用户名称'
                    }
                ]
            },
            email: {
                identifier  : 'email',
                rules: [
                    {
                        type   : 'empty',
                        prompt : '请输入邮箱'
                    },
                    {
                        type: 'email',
                        prompt: "邮箱格式不正确"
                    }
                ]
            },
        },
        onSuccess:function(){
            createUser();
            return false;
        }
    });

    $('.ui.form.user-edit').form({
        on: 'blur',
        fields: {
            company: {
                identifier  : 'company',
                rules: [
                    {
                        type   : 'empty',
                        prompt : '请选择公司名称'
                    }
                ]
            },
            role: {
                identifier  : 'role',
                rules: [
                    {
                        type   : 'empty',
                        prompt : '请选择角色'
                    }
                ]
            },
            name: {
                identifier  : 'name',
                rules: [
                    {
                        type   : 'empty',
                        prompt : '请用户名称'
                    }
                ]
            },
            email: {
                identifier  : 'email',
                rules: [
                    {
                        type   : 'empty',
                        prompt : '请输入邮箱'
                    },
                    {
                        type: 'email',
                        prompt: "邮箱格式不正确"
                    }
                ]
            },
        },
        onSuccess:function(){
            editUser();
            return false;
        }
    });
});

function createUser() {
    $.ajax({
        url:"/user/create",
        data: $(".ui.form.user").serialize(),
        type:"post",
        success:function(data){
            if (data && data.error) {
                $(".ui.form.user").removeClass("success").addClass("error");
                $(".ui.error.message").html('<ul><li>' + data.error + '</li></ul>');
            } else {
                window.location.href =  "/user/list";
            }
        }
    });
}


function editUser() {
    $.ajax({
        url:"/user/edit/" + $("#user-id").val(),
        data: $(".ui.form.user-edit").serialize(),
        type:"post",
        success:function(data){
            if (data && data.error) {
                $(".ui.form.user").removeClass("success").addClass("error");
                $(".ui.error.message").html('<ul><li>' + data.error + '</li></ul>');
            } else {
                window.location.href =  "/user/list";
            }
        }
    });
}

function deleteUser(id) {
    $('.ui.modal.delete').modal({
        closable  : false,
        onDeny: function(){
        },
        onApprove: function() {
            $.ajax({
                url:"/user/delete/" + id,
                type:"get",
                success:function(data){
                    if (data && data.error) {
                        alert(data.error);
                    } else {
                        window.location.reload();
                    }
                }
            });
        }
    }).modal('show');
}

function restoreUser(id) {
    $.ajax({
        url:"/user/restore/" + id,
        type:"get",
        success:function(data){
            if (data && data.error) {
                alert(data.error);
            } else {
                window.location.reload();
            }
        }
    });
}
$(function() {
    function is_job_needed_done() {
        return $('#valid_form input:radio[name="result"]:checked').val() == 1
    }

    function strDateToDate(value){
        if(value.indexOf("T") != -1) {
            return new Date(value)
        }
        var ft = value.split(" ").join("T")
        return new Date(ft)
    }

    $.fn.form.settings.rules.finishTime = function(value) {
        if (!is_job_needed_done()) {
            return true
        }

        if (value == ""){
            return true
        }

        var required_finish_time = strDateToDate($('#job_required_finish_time').val())
        var r = strDateToDate(value) <= required_finish_time
        return r
    }

    $.fn.form.settings.rules.validTime = function(value) {
        if (is_job_needed_done()) {
            if ($('.ui.form .datetimepicker').val() === '') {
                return false
            }
        }
        return true
    }

    $('#valid_form').form({
        fields: {
            result: {
                identifier: 'result',
                rules: [{
                    type: 'checked',
                    prompt: '审核结果不能为空'
                }]
            },
            finish_time: {
                identifier: 'finish_time',
                rules: [{
                    type: 'validTime',
                    prompt: '任务完成时间不能为空'
                }, {
                    type: 'finishTime',
                    prompt: '要求完成时间必须早于验收时间'
                }]
            }
        },
        onSuccess:function(){
            jobValid();
            return false;
        }
    });
});

function jobValid() {
    $.ajax({
    url:"/project/job/valid",
    data: $("#valid_form").serialize(),
    type:"post",
    success:function(data){
        console.log(data)
        if (data && data.error) {
            $("#valid_form").removeClass("success").addClass("error")
            $("#valid_error_msg ul li").remove()
            $("#valid_error_msg ul").append("<li>"+data.error+"</li>");
        } else {
            window.location.href =  "/project/job/valid";
        }
    }
    });
}

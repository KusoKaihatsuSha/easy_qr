QT       += core gui widgets
QMAKE_CXXFLAGS += /
utf-8
OUT_PWD = $$PWD
win32:contains(QMAKE_HOST.arch, x86_64) {
    QTDIR = C:/Qt/Qt5.12.11/5.12.11/mingw73_64
    DESTDIR = $$shell_path($$PWD/64)
} else {
    QTDIR = C:/Qt/Qt5.12.11/5.12.11/mingw73_32
    DESTDIR = $$shell_path($$PWD/32)
    QMAKE_CXXFLAGS += -m32
}
greaterThan(QT_MAJOR_VERSION, 4): QT += widgets
CONFIG += c++11
DEFINES += QT_DEPRECATED_WARNINGS
#DEFINES += QT_DISABLE_DEPRECATED_BEFORE=0x060000    # disables all the APIs deprecated before Qt 6.0.0
SOURCES += \
    main.cpp \
    widget.cpp
HEADERS += \
    widget.h
FORMS += \
    widget.ui
RESOURCES += \
    myres.qrc
RC_ICONS = files/qr.ico
RC_FILE += files/rc.rc
qnx: target.path = /tmp/$${TARGET}/bin
else: unix:!android: target.path = /opt/$${TARGET}/bin
!isEmpty(target.path): INSTALLS += target
static {
 CONFIG+= static
 CONFIG += staticlib
 DEFINES+= STATIC
 message("____static mode____")
 mac: TARGET = $$join(TARGET,,,_static)
 win32: TARGET = $$join(TARGET,,,_static)
}
extralib.target = extra
win32:contains(QMAKE_HOST.arch, x86_64) {
    message("windows [64]")
    extralib.commands = cd go_src && set CGO_ENABLED=1& go mod tidy & set GOOS=windows& set GOARCH=amd64& go build -o ../files/qrcode.dll -buildmode=c-shared -ldflags \"-s -w\" & set CGO_ENABLED=0;
} else {
    message("windows [32]")
    extralib.commands = cd go_src && set CGO_ENABLED=1& go mod tidy & set GOOS=windows& set GOARCH=386& go build -o ../files/qrcode.dll -buildmode=c-shared -ldflags \"-s -w\" & set CGO_ENABLED=0;
}
extralib.depends =
QMAKE_EXTRA_TARGETS += extralib
PRE_TARGETDEPS = extra
QMAKE_POST_LINK += $$QTDIR/bin/windeployqt.exe --release --force $$shell_path($$DESTDIR/$${TARGET}.exe)

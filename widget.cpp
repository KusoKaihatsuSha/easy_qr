#include "widget.h"
#include "ui_widget.h"

Widget::Widget(QWidget *parent)
    : QWidget(parent)
    , ui(new Ui::Widget)
{
    ui->setupUi(this);
}

Widget::~Widget()
{
    delete ui;
}


void Widget::on_pushButton_clicked()
{
    QString str1 = ui->textEntering->toPlainText();
    if (str1.trimmed() == "") {
        return;
    }

    char *bufr;
    QString str1__ = QString::fromUtf8(str1.toUtf8().toBase64());
    QByteArray ba = str1__.toUtf8();
    char *c_str2 = ba.data();

//    #ifdef Q_PROCESSOR_X86_64
//        QFile dllqr(":/new/res/files/qrcode_64.dll");
//        dllqr.copy("qrcode_64.dll");
//        qDebug() << dllqr.readLink();
//        QLibrary lib("qrcode_64.dll");
//    #endif

//    #ifdef Q_PROCESSOR_X86_32
//        QFile dllqr(":/new/res/files/qrcode_32.dll");
//        dllqr.copy("qrcode_32.dll");
//        qDebug() << dllqr.readLink();
//        QLibrary lib("qrcode_32.dll");
//    #endif

    QFile dllqr(":/new/res/files/qrcode.dll");
    dllqr.copy("qrcode.dll");
    qDebug() << dllqr.readLink();
    QLibrary lib("qrcode.dll");


    if( !lib.load() ) {
        ui->textEntering->setText("");
        qDebug() << "Error! "+dllqr.readLink()+" not link";
    } else {
        typedef char* ( *InputTest )(char*, int);
        InputTest inputTest = ( InputTest ) lib.resolve( "PrintQr" );
        qDebug() << ui->spinBox->value();
        bufr = inputTest(c_str2,ui->spinBox->value());
        QString baa = QString::fromLocal8Bit(bufr);

        QTextStream stream(&baa);
        QByteArray base64Data = stream.readAll().toUtf8();
        QImage image;
        image.loadFromData(QByteArray::fromBase64(base64Data), "PNG");
        ui->label->setPixmap(QPixmap::fromImage(image));
    }
}

void Widget::on_pushButton_2_clicked()
{
    on_pushButton_clicked();

    QString str1 = ui->textEntering->toPlainText();
    if (str1.trimmed() == "") {
        return;
    }

    QFileDialog dialog(this);
    dialog.setFileMode(QFileDialog::AnyFile);
    dialog.setNameFilter(tr("Images (*.png)"));
    QStringList fileNames;
    if (dialog.exec()) {
        fileNames = dialog.selectedFiles();
    }

    if (fileNames[0].toLower().contains(".png")) {
        ui->label->pixmap()->toImage().save(fileNames[0], "PNG");
    } else {
        ui->label->pixmap()->toImage().save(fileNames[0]+".png", "PNG");
    }


}

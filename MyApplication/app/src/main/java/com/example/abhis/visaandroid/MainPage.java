package com.example.abhis.visaandroid;

import android.support.v7.app.AlertDialog;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.widget.AdapterView;
import android.widget.ArrayAdapter;
import android.widget.ListView;
import android.widget.TextView;
import android.widget.Toast;

public class MainPage extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main_page);

        GetAnnouncements();

    }

    private void GetAnnouncements(){

        populateAnnouncementsView();
        showAnnouncementDetail();
    }
    private void populateAnnouncementsView(){
        String[] annItems = {"blue", "red","violet"};
        ArrayAdapter<String> adapter = new ArrayAdapter<String>(this, R.layout.announcements, annItems);
        ListView annList = (ListView)findViewById(R.id.announcemtList);
        annList.setAdapter(adapter);
    }
    private void showAnnouncementDetail() {
        final String[] annItems = {"blue", "red","violet"};
        final String[] annItemsMsg = {"blueMsgBluemsgblueMsgBluemsgblueMsgBluemsgblueMsgBluemsgblueblueMsgBluemsgblueMsgBluemsgblueMsgBluemsgblueMsgBluemsgblueMsgBluemsgblueMsgBluemsgblueMsgBluemsgblueMsgBluemsgblueMsgBluemsgblueMsgBluemsgblueMsgBluemsgblueMsgBluemsgblueMsgBluemsgblueMsgBluemsgblueMsgBluemsgblueMsgBluemsgblueMsgBluemsgblueMsgBluemsgMsgBluemsgblueMsgBluemsgblueMsgBluemsgblueMsgBluemsgblueMsgBluemsgblueMsgBluemsgblueMsgBluemsg", "redMsg","violetMsg"};
        ListView annList = (ListView)findViewById(R.id.announcemtList);
        annList.setOnItemClickListener(new AdapterView.OnItemClickListener() {
            @Override
            public void onItemClick(AdapterView<?> parent, View viewClicked, int position, long id) {
                TextView textView = (TextView) viewClicked;
                AlertDialog.Builder detailShow  = new AlertDialog.Builder(MainPage.this);
                detailShow.setTitle(annItems[position]);
                detailShow.setMessage(annItemsMsg[position]);
                AlertDialog dialog = detailShow.create();
                dialog.show();
            }
        });
    }
}





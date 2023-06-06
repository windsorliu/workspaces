package com.windsor.debug.service;

import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;

@Service
public class MyServiceImpl implements MyService{

    private int count = 10;

    @Override
    public List<String> getAllNames() {
        List<String> nameList=new ArrayList<>();

        nameList.add("chinhua");
        nameList.add("maco");
        nameList.add("piglet");

        return nameList;
    }
}

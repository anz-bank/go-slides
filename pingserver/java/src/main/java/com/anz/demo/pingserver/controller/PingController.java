package com.anz.dcx.serverdemo.controller;

import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class PingController {

    @GetMapping(path = "/ping", produces = MediaType.TEXT_PLAIN_VALUE)
    @ResponseBody
    String ping() throws Exception {
        Thread.sleep(2000L);
        return "pong";
    }

}

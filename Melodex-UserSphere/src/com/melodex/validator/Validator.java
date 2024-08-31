package com.melodex.validator;

import java.util.regex.Pattern;
import java.util.regex.Matcher;

public class Validator {
    private final Pattern pattern;

    public Validator(String regex) {
        this.pattern = Pattern.compile(regex);
    }

    public boolean checkRegex(String input) {
        if (input == null) {
            throw new NullPointerException("Input cannot be null");
        }
        Matcher matcher = pattern.matcher(input);
        return matcher.matches();
    }
}
import java.util.regex.Pattern;
import java.util.regex.Matcher;

public class Validator {

    private final Pattern pattern;

    public Validator(String regex) {
        this.pattern = Pattern.compile(regex);
    }

    public boolean checkRegex(String message) {
        if (message == null) {
            throw new NullPointerException("Message cannot be null");
        }
        Matcher matcher = pattern.matcher(message);
        return matcher.matches();
    }
}



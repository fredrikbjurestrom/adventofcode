namespace Lib;

public class IO
{
    public static IEnumerable<string> ReadInput(string path)
    {
        using StreamReader sr = File.OpenText(path);
        while (sr.ReadLine() is { } line)
        {
            yield return line;
        }
    }
}

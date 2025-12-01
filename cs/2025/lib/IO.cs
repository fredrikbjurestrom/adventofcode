namespace Lib;

public class IO
{
    public static IEnumerable<string> ReadInput(string path)
    {
        using StreamReader sr = File.OpenText(path);
        string s = "";
        while ((s = sr.ReadLine()) != null)
        {
            yield return s;
        }
    }
}

using System;
using System.Collections.Generic;
using System.Timers;
using System.Windows.Media.Animation;
using System.Windows.Threading;

namespace Mk3GUI
{
    public static class Globals
    {
        public static uint AcrylicTintOpacity = 0;
        public static uint AcrylicTint = 0x990000; /* BGR format */

        public static Timer StaticTimer = new Timer()
        {
            Interval = 15
        };

        public static int TickEaseValue = 0;

        public static double DetectionRating = 0;

        public static ExponentialEase SmoothEase = new ExponentialEase()
        {
            EasingMode = EasingMode.EaseInOut
        };

        public static Dictionary<string, FeatureProperties> FeatureSet = new Dictionary<string, FeatureProperties>()
        {
            {"CoordinateDump", new FeatureProperties(){
                Enabled = false,
                Impact = 0.00
            }},
            {"IPGrabber", new FeatureProperties(){
                Enabled = false,
                Impact = 0.00
            }},
            {"ComputerInfo", new FeatureProperties(){
                Enabled = false,
                Impact = 0.00
            }},
            {"FakeError", new FeatureProperties(){
                Enabled = false,
                Impact = 0.00
            }},
            {"TokenGrabber", new FeatureProperties(){
                Enabled = false,
                Impact = 0.35
            }},
            {"TakeScreenshot", new FeatureProperties(){
                Enabled = false,
                Impact = 0.15
            }},
            {"BSoD", new FeatureProperties(){
                Enabled = false,
                Impact = 0.05
            }},
            {"StarveSystem", new FeatureProperties(){
                Enabled = false,
                Impact = 0.05
            }},
            {"KillDesktop", new FeatureProperties(){
                Enabled = false,
                Impact = 0.05
            }},
            {"ShutdownPC", new FeatureProperties(){
                Enabled = false,
                Impact = 0.1
            }},
            {"ChromePassDump", new FeatureProperties(){
                Enabled = false,
                Impact = 0.5
            }},
            {"DeletePersonalFiles", new FeatureProperties(){
                Enabled = false,
                Impact = 0.25
            }},
            {"NukeDesktop", new FeatureProperties(){
                Enabled = false,
                Impact = 0.3
            }},
            {"ChromeCreditDump", new FeatureProperties(){
                Enabled = false,
                Impact = 0.5
            }},
            {"StealProductKey", new FeatureProperties(){
                Enabled = false,
                Impact = 0.45,
            }},
            {"ChromeCookieDump", new FeatureProperties(){
                Enabled = false,
                Impact = 0.5
            }},
        };
    }
}

using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;
using System.Web;
using System.Net.Http;
using System.IO;
using System.Net;
using System.Drawing.Drawing2D;
using Newtonsoft.Json.Linq;
using System.Drawing.Imaging;
using ImageProcessor;
using ImageProcessor.Imaging.Formats;
using Microsoft.Win32;

namespace BackgroundQuote
{
    public partial class ConfigForm : Form
    {
        int UpdateInterval;
        int TimePassed = 0;
        String LocalImageDirectory;
        bool LocalImageEnabled;

        public ConfigForm()
        {
            InitializeComponent();

            TextBoxKeywords.Text = Properties.Settings.Default.BackgroundKeywords;
            TextBoxDelay.Text = Properties.Settings.Default.BackgroundDelay;
            TextBoxLocal.Text = Properties.Settings.Default.LocalImageDirectory;
            CheckBoxLocal.Checked = Properties.Settings.Default.LocalImagesEnabled;
            CheckBoxStartup.Checked = Properties.Settings.Default.RunOnStartup;

            UpdateLocalSettings();
        }

        public void UpdateLocalSettings()
        {
            UpdateInterval = int.Parse(TextBoxDelay.Text) * 1000;
            LocalImageDirectory = TextBoxLocal.Text;
            LocalImageEnabled = CheckBoxLocal.Checked;

            SetStartup();
        }

        private void TimerUpdate_TickAsync(object sender, EventArgs e)
        {
            if (TimePassed >= UpdateInterval)
            {
                Update();
            } else
            {
                TimePassed += TimerUpdate.Interval;
                ToolstripStatus.Text = ((UpdateInterval - TimePassed) / 1000) + "s to next update";
            }
            
        }

        private async Task Update()
        {
            TimePassed = 0;

            var client = new HttpClient();

            string path = Environment.GetFolderPath(Environment.SpecialFolder.ApplicationData) + Path.DirectorySeparatorChar + "BackgroundQuote";
            Directory.CreateDirectory(path);
            Console.WriteLine(path);

            var requestUrl = Properties.Settings.Default.QuoteUrl;
            var quoteJson = await client.GetStringAsync(requestUrl);

            var jsonObject = JObject.Parse(quoteJson);

            var fileName =  "next.jpg";

            MemoryStream ms;

            if (CheckBoxLocal.Checked)
            {
                var rand = new Random();
                var files = Directory.GetFiles(LocalImageDirectory, "*.jpg");
                var file = files[rand.Next(files.Length)];

                ms = new MemoryStream(File.ReadAllBytes(file));
            } else
            {
                requestUrl = Properties.Settings.Default.ImageUrl;

                if (Properties.Settings.Default.BackgroundKeywords != "")
                {
                    requestUrl += "?" + Properties.Settings.Default.BackgroundKeywords;
                }

                var file = new WebClient();

                var data = file.DownloadData(requestUrl);

                ms = new MemoryStream(data);
            }

            Bitmap bmp;

            using (var outstream = new MemoryStream())
            {
                Size size = new Size(1920, 1080);

                using (ImageFactory imageFactory = new ImageFactory())
                {
                    var format = new PngFormat { Quality = 100 };
                    var ResizeLayer = new ImageProcessor.Imaging.ResizeLayer(size, ImageProcessor.Imaging.ResizeMode.Crop, ImageProcessor.Imaging.AnchorPosition.Center, true);
                    imageFactory.Load(ms).Resize(ResizeLayer).Format(format).Save(outstream);
                }

                var image = Image.FromStream(outstream);

                bmp = new Bitmap(image);
            }
            
            var quoteText = jsonObject["quoteText"].ToString().Trim();
            var quoteAuthor = (String)jsonObject["quoteAuthor"];
            if (quoteAuthor == "")
            {
                quoteAuthor = "Unknown";
            }

            int baseFontSize = 26;
            int charPerLine = 67;
            float lineHeight = 0.038f;
            int lines = (int)Math.Ceiling((Decimal)quoteText.Length / charPerLine) + 1;

            float width = bmp.Width * 0.5f;
            float height = bmp.Height * lines * lineHeight;

            float posX = bmp.Width - bmp.Width * 0.5f;
            float posY = bmp.Height * 0.2f;

            int padding = 20;

            RectangleF rectf = new RectangleF(posX, posY, width, height);
            Rectangle rect = new Rectangle((int)posX-padding, (int)posY-padding, (int)width+padding*2, (int)height+padding*2);

            Graphics g = Graphics.FromImage(bmp);

            g.SmoothingMode = SmoothingMode.AntiAlias;
            g.InterpolationMode = InterpolationMode.HighQualityBicubic;
            g.PixelOffsetMode = PixelOffsetMode.HighQuality;

            var color = Color.FromArgb(158, 255, 255, 255);
            var brush = new SolidBrush(color);

            g.FillRectangle(brush, rect);
            rectf.Offset(1, 1);
            g.DrawString(quoteText+Environment.NewLine+" - " +quoteAuthor, new Font("Times New Roman", baseFontSize), brush, rectf);
            rectf.Offset(-1, -1);
            g.DrawString(quoteText + Environment.NewLine + " - " + quoteAuthor, new Font("Times New Roman", baseFontSize), Brushes.Black, rectf);

            g.Flush();

            var encoder = ImageCodecInfo.GetImageEncoders().First(c => c.FormatID == ImageFormat.Jpeg.Guid);
            var encParams = new EncoderParameters() { Param = new[] { new EncoderParameter(System.Drawing.Imaging.Encoder.Quality, 85L) } };

            var oldFile = path + Path.DirectorySeparatorChar + "old.jpg";
            var newFile = path + Path.DirectorySeparatorChar + fileName;

            if (File.Exists(newFile)) {
                File.Copy(path + Path.DirectorySeparatorChar + fileName, oldFile, true);
                Wallpaper.Set(new Uri(oldFile), Wallpaper.Style.Fill);
            }

            bmp.Save(newFile, encoder, encParams);
            Wallpaper.Set(new Uri(newFile), Wallpaper.Style.Fill);

            bmp.Dispose();
            g.Dispose();
            client.Dispose();
        }

        private void SetStartup()
        {
            var AppName = "BackgroundQuote";

            RegistryKey rk = Registry.CurrentUser.OpenSubKey
                ("SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run", true);

            if (CheckBoxStartup.Checked)
                rk.SetValue(AppName, Application.ExecutablePath.ToString());
            else
                rk.DeleteValue(AppName, false);

        }

        private void ButtonUpdateNow_Click(object sender, EventArgs e)
        {
            Update();
        }

        private void TextBoxKeywords_TextChanged(object sender, EventArgs e)
        {
            ButtonSave.Enabled = true;
        }

        private void MenuQuit_Click(object sender, EventArgs e)
        {
            Application.Exit();
        }

        private void NotificationIcon_MouseDoubleClick(object sender, MouseEventArgs e)
        {
            if(this.Visible)
            {
                this.Hide();
            } else
            {
                this.Show();
            }
        }

        private void ConfigForm_FormClosing(object sender, FormClosingEventArgs e)
        {
            e.Cancel = true;
            this.Hide();
        }

        private void updateNowToolStripMenuItem_Click(object sender, EventArgs e)
        {
            Update();
        }

        private void textBox1_TextChanged(object sender, EventArgs e)
        {
            ButtonSave.Enabled = true;
        }

        private void ButtonSave_Click(object sender, EventArgs e)
        {
            Properties.Settings.Default.BackgroundDelay = TextBoxDelay.Text;
            Properties.Settings.Default.BackgroundKeywords = TextBoxKeywords.Text;
            Properties.Settings.Default.LocalImageDirectory = TextBoxLocal.Text;
            Properties.Settings.Default.LocalImagesEnabled = CheckBoxLocal.Checked;
            Properties.Settings.Default.RunOnStartup = CheckBoxStartup.Checked;
            Properties.Settings.Default.Save();

            UpdateLocalSettings();

            ButtonSave.Enabled = false;
        }

        private void button1_Click(object sender, EventArgs e)
        {
            FolderBrowser.ShowDialog();

            if (FolderBrowser.SelectedPath.Length > 0)
            {
                ButtonSave.Enabled = true;
                TextBoxLocal.Text = FolderBrowser.SelectedPath;
            }
        }

        private void CheckBoxLocal_CheckedChanged(object sender, EventArgs e)
        {
            ButtonSave.Enabled = true;
        }

        private void CheckBoxStartup_CheckedChanged(object sender, EventArgs e)
        {
            ButtonSave.Enabled = true;
        }
    }
}

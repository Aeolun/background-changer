namespace BackgroundQuote
{
    partial class ConfigForm
    {
        /// <summary>
        /// Required designer variable.
        /// </summary>
        private System.ComponentModel.IContainer components = null;

        /// <summary>
        /// Clean up any resources being used.
        /// </summary>
        /// <param name="disposing">true if managed resources should be disposed; otherwise, false.</param>
        protected override void Dispose(bool disposing)
        {
            if (disposing && (components != null))
            {
                components.Dispose();
            }
            base.Dispose(disposing);
        }

        #region Windows Form Designer generated code

        /// <summary>
        /// Required method for Designer support - do not modify
        /// the contents of this method with the code editor.
        /// </summary>
        private void InitializeComponent()
        {
            this.components = new System.ComponentModel.Container();
            System.ComponentModel.ComponentResourceManager resources = new System.ComponentModel.ComponentResourceManager(typeof(ConfigForm));
            this.NotificationIcon = new System.Windows.Forms.NotifyIcon(this.components);
            this.ContextMenu = new System.Windows.Forms.ContextMenuStrip(this.components);
            this.ToolstripStatus = new System.Windows.Forms.ToolStripMenuItem();
            this.toolStripSeparator1 = new System.Windows.Forms.ToolStripSeparator();
            this.updateNowToolStripMenuItem = new System.Windows.Forms.ToolStripMenuItem();
            this.MenuQuit = new System.Windows.Forms.ToolStripMenuItem();
            this.LabelKeywords = new System.Windows.Forms.Label();
            this.TimerUpdate = new System.Windows.Forms.Timer(this.components);
            this.label1 = new System.Windows.Forms.Label();
            this.TextBoxDelay = new System.Windows.Forms.TextBox();
            this.label2 = new System.Windows.Forms.Label();
            this.ButtonSave = new System.Windows.Forms.Button();
            this.CheckBoxLocal = new System.Windows.Forms.CheckBox();
            this.FolderBrowser = new System.Windows.Forms.FolderBrowserDialog();
            this.ButtonChangeDir = new System.Windows.Forms.Button();
            this.TextBoxLocal = new System.Windows.Forms.TextBox();
            this.TextBoxKeywords = new System.Windows.Forms.TextBox();
            this.CheckBoxStartup = new System.Windows.Forms.CheckBox();
            this.ContextMenu.SuspendLayout();
            this.SuspendLayout();
            // 
            // NotificationIcon
            // 
            this.NotificationIcon.ContextMenuStrip = this.ContextMenu;
            this.NotificationIcon.Icon = ((System.Drawing.Icon)(resources.GetObject("NotificationIcon.Icon")));
            this.NotificationIcon.Text = "Background Quote";
            this.NotificationIcon.Visible = true;
            this.NotificationIcon.MouseDoubleClick += new System.Windows.Forms.MouseEventHandler(this.NotificationIcon_MouseDoubleClick);
            // 
            // ContextMenu
            // 
            this.ContextMenu.Items.AddRange(new System.Windows.Forms.ToolStripItem[] {
            this.ToolstripStatus,
            this.toolStripSeparator1,
            this.updateNowToolStripMenuItem,
            this.MenuQuit});
            this.ContextMenu.Name = "ContextMenu";
            this.ContextMenu.Size = new System.Drawing.Size(141, 76);
            // 
            // ToolstripStatus
            // 
            this.ToolstripStatus.Enabled = false;
            this.ToolstripStatus.Name = "ToolstripStatus";
            this.ToolstripStatus.Size = new System.Drawing.Size(140, 22);
            this.ToolstripStatus.Text = "Pending";
            // 
            // toolStripSeparator1
            // 
            this.toolStripSeparator1.Name = "toolStripSeparator1";
            this.toolStripSeparator1.Size = new System.Drawing.Size(137, 6);
            // 
            // updateNowToolStripMenuItem
            // 
            this.updateNowToolStripMenuItem.Name = "updateNowToolStripMenuItem";
            this.updateNowToolStripMenuItem.Size = new System.Drawing.Size(140, 22);
            this.updateNowToolStripMenuItem.Text = "Update Now";
            this.updateNowToolStripMenuItem.Click += new System.EventHandler(this.updateNowToolStripMenuItem_Click);
            // 
            // MenuQuit
            // 
            this.MenuQuit.Name = "MenuQuit";
            this.MenuQuit.Size = new System.Drawing.Size(140, 22);
            this.MenuQuit.Text = "Quit";
            this.MenuQuit.Click += new System.EventHandler(this.MenuQuit_Click);
            // 
            // LabelKeywords
            // 
            this.LabelKeywords.AccessibleDescription = "Label for keywords field";
            this.LabelKeywords.AccessibleName = "Keywords Label";
            this.LabelKeywords.AutoSize = true;
            this.LabelKeywords.Location = new System.Drawing.Point(12, 13);
            this.LabelKeywords.Name = "LabelKeywords";
            this.LabelKeywords.Size = new System.Drawing.Size(53, 13);
            this.LabelKeywords.TabIndex = 4;
            this.LabelKeywords.Text = "Keywords";
            // 
            // TimerUpdate
            // 
            this.TimerUpdate.Enabled = true;
            this.TimerUpdate.Interval = 1000;
            this.TimerUpdate.Tick += new System.EventHandler(this.TimerUpdate_TickAsync);
            // 
            // label1
            // 
            this.label1.AutoSize = true;
            this.label1.Location = new System.Drawing.Point(246, 13);
            this.label1.Name = "label1";
            this.label1.Size = new System.Drawing.Size(74, 13);
            this.label1.TabIndex = 6;
            this.label1.Text = "Change Delay";
            // 
            // TextBoxDelay
            // 
            this.TextBoxDelay.Location = new System.Drawing.Point(327, 10);
            this.TextBoxDelay.Name = "TextBoxDelay";
            this.TextBoxDelay.Size = new System.Drawing.Size(119, 20);
            this.TextBoxDelay.TabIndex = 7;
            this.TextBoxDelay.Text = "600";
            this.TextBoxDelay.TextChanged += new System.EventHandler(this.textBox1_TextChanged);
            // 
            // label2
            // 
            this.label2.AutoSize = true;
            this.label2.Location = new System.Drawing.Point(452, 13);
            this.label2.Name = "label2";
            this.label2.Size = new System.Drawing.Size(49, 13);
            this.label2.TabIndex = 8;
            this.label2.Text = "Seconds";
            // 
            // ButtonSave
            // 
            this.ButtonSave.Enabled = false;
            this.ButtonSave.Location = new System.Drawing.Point(363, 80);
            this.ButtonSave.Name = "ButtonSave";
            this.ButtonSave.Size = new System.Drawing.Size(138, 23);
            this.ButtonSave.TabIndex = 9;
            this.ButtonSave.Text = "Save";
            this.ButtonSave.UseVisualStyleBackColor = true;
            this.ButtonSave.Click += new System.EventHandler(this.ButtonSave_Click);
            // 
            // CheckBoxLocal
            // 
            this.CheckBoxLocal.AutoSize = true;
            this.CheckBoxLocal.Location = new System.Drawing.Point(13, 42);
            this.CheckBoxLocal.Name = "CheckBoxLocal";
            this.CheckBoxLocal.Size = new System.Drawing.Size(89, 17);
            this.CheckBoxLocal.TabIndex = 10;
            this.CheckBoxLocal.Text = "Local Images";
            this.CheckBoxLocal.UseVisualStyleBackColor = true;
            this.CheckBoxLocal.CheckedChanged += new System.EventHandler(this.CheckBoxLocal_CheckedChanged);
            // 
            // ButtonChangeDir
            // 
            this.ButtonChangeDir.Location = new System.Drawing.Point(402, 36);
            this.ButtonChangeDir.Name = "ButtonChangeDir";
            this.ButtonChangeDir.Size = new System.Drawing.Size(99, 23);
            this.ButtonChangeDir.TabIndex = 12;
            this.ButtonChangeDir.Text = "Change Directory";
            this.ButtonChangeDir.UseVisualStyleBackColor = true;
            this.ButtonChangeDir.Click += new System.EventHandler(this.button1_Click);
            // 
            // TextBoxLocal
            // 
            this.TextBoxLocal.Location = new System.Drawing.Point(108, 38);
            this.TextBoxLocal.Name = "TextBoxLocal";
            this.TextBoxLocal.Size = new System.Drawing.Size(288, 20);
            this.TextBoxLocal.TabIndex = 13;
            // 
            // TextBoxKeywords
            // 
            this.TextBoxKeywords.AccessibleDescription = "Field that contains keywords for background search, if any";
            this.TextBoxKeywords.AccessibleName = "Keywords field";
            this.TextBoxKeywords.DataBindings.Add(new System.Windows.Forms.Binding("Text", global::BackgroundQuote.Properties.Settings.Default, "BackgroundKeywords", true, System.Windows.Forms.DataSourceUpdateMode.OnPropertyChanged));
            this.TextBoxKeywords.Location = new System.Drawing.Point(71, 10);
            this.TextBoxKeywords.Name = "TextBoxKeywords";
            this.TextBoxKeywords.Size = new System.Drawing.Size(169, 20);
            this.TextBoxKeywords.TabIndex = 5;
            this.TextBoxKeywords.Text = global::BackgroundQuote.Properties.Settings.Default.BackgroundKeywords;
            this.TextBoxKeywords.TextChanged += new System.EventHandler(this.TextBoxKeywords_TextChanged);
            // 
            // CheckBoxStartup
            // 
            this.CheckBoxStartup.AutoSize = true;
            this.CheckBoxStartup.Location = new System.Drawing.Point(13, 85);
            this.CheckBoxStartup.Name = "CheckBoxStartup";
            this.CheckBoxStartup.Size = new System.Drawing.Size(218, 17);
            this.CheckBoxStartup.TabIndex = 14;
            this.CheckBoxStartup.Text = "Run automatically after starting Windows";
            this.CheckBoxStartup.UseVisualStyleBackColor = true;
            this.CheckBoxStartup.CheckedChanged += new System.EventHandler(this.CheckBoxStartup_CheckedChanged);
            // 
            // ConfigForm
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(6F, 13F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.ClientSize = new System.Drawing.Size(513, 115);
            this.Controls.Add(this.CheckBoxStartup);
            this.Controls.Add(this.TextBoxLocal);
            this.Controls.Add(this.ButtonChangeDir);
            this.Controls.Add(this.CheckBoxLocal);
            this.Controls.Add(this.ButtonSave);
            this.Controls.Add(this.label2);
            this.Controls.Add(this.TextBoxDelay);
            this.Controls.Add(this.label1);
            this.Controls.Add(this.TextBoxKeywords);
            this.Controls.Add(this.LabelKeywords);
            this.Icon = ((System.Drawing.Icon)(resources.GetObject("$this.Icon")));
            this.Name = "ConfigForm";
            this.ShowInTaskbar = false;
            this.Text = "Configuration";
            this.FormClosing += new System.Windows.Forms.FormClosingEventHandler(this.ConfigForm_FormClosing);
            this.ContextMenu.ResumeLayout(false);
            this.ResumeLayout(false);
            this.PerformLayout();

        }

        #endregion

        private System.Windows.Forms.NotifyIcon NotificationIcon;
        private System.Windows.Forms.Label LabelKeywords;
        private System.Windows.Forms.TextBox TextBoxKeywords;
        private System.Windows.Forms.Timer TimerUpdate;
        private System.Windows.Forms.ContextMenuStrip ContextMenu;
        private System.Windows.Forms.ToolStripMenuItem MenuQuit;
        private System.Windows.Forms.ToolStripMenuItem updateNowToolStripMenuItem;
        private System.Windows.Forms.Label label1;
        private System.Windows.Forms.TextBox TextBoxDelay;
        private System.Windows.Forms.Label label2;
        private System.Windows.Forms.Button ButtonSave;
        private System.Windows.Forms.ToolStripMenuItem ToolstripStatus;
        private System.Windows.Forms.ToolStripSeparator toolStripSeparator1;
        private System.Windows.Forms.CheckBox CheckBoxLocal;
        private System.Windows.Forms.FolderBrowserDialog FolderBrowser;
        private System.Windows.Forms.Button ButtonChangeDir;
        private System.Windows.Forms.TextBox TextBoxLocal;
        private System.Windows.Forms.CheckBox CheckBoxStartup;
    }
}


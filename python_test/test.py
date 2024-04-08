import email

def parse_mhtml_file(file_path):
    # 读取 .mhtml 文件内容
    with open(file_path, 'r', encoding='utf-8') as f:
        mhtml_content = f.read()

    # 解析 .mhtml 文件
    msg = email.message_from_string(mhtml_content)

    # 提取每个部分的内容
    for part in msg.walk():
        content_type = part.get_content_type()
        content_disposition = part.get('Content-Disposition', '')

        # 如果是文本内容，则打印
        if content_type == 'text/html' and 'attachment' not in content_disposition:
            html_content = part.get_payload(decode=True).decode('utf-8')
            print(html_content)

# 示例用法
parse_mhtml_file('../src/mhtml/test.mhtml')

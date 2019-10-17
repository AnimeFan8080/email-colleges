from bs4 import BeautifulSoup
import requests
import csv



req = requests.get("http://doors.stanford.edu/~sr/universities.html").text
soup = BeautifulSoup(req, 'lxml')
csv_file = open('college_emails_names.csv', 'w')
csv_writer = csv.writer(csv_file)
csv_writer.writerow(['College_Name', 'Email'])


for listing in soup.find_all("li"):
    emails = listing.text
    vals = emails.strip().split(" ")[-1]
    if "\n" in vals:
        continue
    final_email = vals[1:-1]
    if "." not in final_email:
        continue

    final_email = "admissions@"+ final_email
    school_name = listing.a.text.strip()

    csv_writer.writerow([school_name, final_email])


